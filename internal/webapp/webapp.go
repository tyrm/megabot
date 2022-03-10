package webapp

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/markbates/pkger"
	"github.com/rbcervilla/redisstore/v8"
	"github.com/spf13/viper"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tyrm/go-util/pkgerutil"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/kv"
	"github.com/tyrm/megabot/internal/kv/redis"
	"github.com/tyrm/megabot/internal/language"
	"github.com/tyrm/megabot/internal/models"
	"github.com/tyrm/megabot/internal/web"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

var tmplFuncs = template.FuncMap{
	"dec": func(i int) int {
		i--
		return i
	},
	"groupSuperAdmin": func() uuid.UUID {
		return models.GroupSuperAdmin()
	},
	"htmlSafe": func(html string) template.HTML {
		/* #nosec G203 */
		return template.HTML(html)
	},
	"inc": func(i int) int {
		i++
		return i
	},
}

// Module contains a webapp module for the web server. Implements web.Module
type Module struct {
	db        db.DB
	store     sessions.Store
	language  *language.Module
	minify    *minify.M
	templates *template.Template

	headLinks     []templateHeadLink
	footerScripts []templateScript

	sigCache     map[string]string
	sigCacheLock sync.RWMutex
}

// New returns a new webapp module
func New(ctx context.Context, db db.DB, r *redis.Client, lMod *language.Module) (web.Module, error) {
	l := logger.WithField("func", "New")

	// Load Templates
	t, err := pkgerutil.CompileTemplates(pkger.Include("/web/template"), "", &tmplFuncs)
	if err != nil {
		return nil, err
	}

	// Fetch new store.
	store, err := redisstore.NewRedisStore(ctx, r.RedisClient())
	if err != nil {
		l.Errorf("create redis store: %s", err.Error())
		return nil, err
	}

	store.KeyPrefix(kv.KeySession())
	store.Options(sessions.Options{
		Path:   pathBase,
		Domain: viper.GetString(config.Keys.ServerExternalHostname),
		MaxAge: 86400 * 60,
	})

	// Register models for GOB
	gob.Register(models.User{})

	// minify
	var m *minify.M
	if viper.GetBool(config.Keys.ServerMinifyHTML) {
		m = minify.New()
		m.AddFunc("text/html", html.Minify)
	}

	// generate head links
	var hl []templateHeadLink
	paths := []string{
		pathFileBootstrapCSS,
		pathFileFontAwesome,
	}
	for _, path := range paths {
		signature, err := getSignature(fmt.Sprintf("%s/%s", staticDir, path))
		if err != nil {
			l.Errorf("getting signature for %s: %s", path, err.Error())
		}

		hl = append(hl, templateHeadLink{
			HRef:        fmt.Sprintf("%s%s", pathStatic, path),
			Rel:         "stylesheet",
			CrossOrigin: "anonymous",
			Integrity:   signature,
		})
	}

	// generate head links
	var fs []templateScript
	scriptPaths := []string{
		pathFileBootstrapJS,
	}
	for _, path := range scriptPaths {
		signature, err := getSignature(fmt.Sprintf("%s/%s", staticDir, path))
		if err != nil {
			l.Errorf("getting signature for %s: %s", path, err.Error())
		}

		fs = append(fs, templateScript{
			Src:         fmt.Sprintf("%s%s", pathStatic, path),
			CrossOrigin: "anonymous",
			Integrity:   signature,
		})
	}

	return &Module{
		db:        db,
		language:  lMod,
		minify:    m,
		templates: t,
		store:     store,

		headLinks:     hl,
		footerScripts: fs,

		sigCache: map[string]string{},
	}, nil
}

// Name return the module name
func (m *Module) Name() string {
	return config.ServerRoleWebapp
}

// Route attaches routes to the web server
func (m *Module) Route(s *web.Server) error {
	// Static Files
	s.PathPrefix(pathStatic + "/").Handler(http.StripPrefix(
		pathStatic+"/", http.FileServer(pkger.Dir(staticDir))))

	webapp := s.PathPrefix(pathBase + "/").Subrouter()
	webapp.Use(m.Middleware)
	webapp.NotFoundHandler = m.notFoundHandler()
	webapp.MethodNotAllowedHandler = m.methodNotAllowedHandler()

	webapp.HandleFunc(pathLogin, m.LoginGetHandler).Methods("GET")
	webapp.HandleFunc(pathLogin, m.LoginPostHandler).Methods("POST")
	webapp.HandleFunc(pathLogout, m.LogoutGetHandler).Methods("GET")

	// Protected Pages
	protected := webapp.PathPrefix("/").Subrouter()
	protected.Use(m.MiddlewareRequireAuth)
	protected.NotFoundHandler = m.notFoundHandler()
	protected.MethodNotAllowedHandler = m.methodNotAllowedHandler()

	protected.HandleFunc(pathHome, m.HomeGetHandler).Methods("GET")
	return nil
}

func (m *Module) getSignatureCached(path string) (string, error) {
	if sig, ok := m.readCachedSignature(path); ok {
		return sig, nil
	}
	sig, err := getSignature(path)
	if err != nil {
		return "", err
	}
	m.writeCachedSignature(path, sig)
	return sig, nil
}

func (m *Module) readCachedSignature(path string) (string, bool) {
	m.sigCacheLock.RLock()
	defer m.sigCacheLock.RUnlock()

	val, ok := m.sigCache[path]
	return val, ok
}

func (m *Module) writeCachedSignature(path string, sig string) {
	m.sigCacheLock.Lock()
	defer m.sigCacheLock.Unlock()

	m.sigCache[path] = sig
	return
}

func getSignature(path string) (string, error) {
	l := logger.WithField("func", "getSignature")
	l.Debugf("getting signature for %s", path)

	file, err := pkger.Open(path)
	if err != nil {
		l.Errorf("opening file: %s", err.Error())
		return "", err
	}

	// read it
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	// hash it
	h := sha512.New384()
	h.Write(data)
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf("sha384-%s", signature), nil
}
