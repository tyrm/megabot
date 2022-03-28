package webapp

import (
	"context"
	"crypto/sha512"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	"github.com/spf13/viper"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tyrm/megabot"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/db"
	"github.com/tyrm/megabot/internal/kv"
	"github.com/tyrm/megabot/internal/kv/redis"
	"github.com/tyrm/megabot/internal/language"
	"github.com/tyrm/megabot/internal/models"
	"github.com/tyrm/megabot/internal/web"
	mbtemplate "github.com/tyrm/megabot/internal/web/template"
	"io/fs"
	"io/ioutil"
	"net/http"
	"sync"
)

// Module contains a webapp module for the web server. Implements web.Module
type Module struct {
	db       db.DB
	store    sessions.Store
	language *language.Module
	minify   *minify.M

	headLinks     []mbtemplate.HeadLink
	footerScripts []mbtemplate.Script

	sigCache     map[string]string
	sigCacheLock sync.RWMutex
}

// New returns a new webapp module
func New(ctx context.Context, db db.DB, r *redis.Client, lMod *language.Module) (web.Module, error) {
	l := logger.WithField("func", "New")

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
	var hl []mbtemplate.HeadLink
	paths := []string{
		pathFileBootstrapCSS,
		pathFileFontAwesome,
		pathFileDefaultCSS,
	}
	for _, path := range paths {
		filePath := staticDir + path
		signature, err := getSignature(filePath)
		if err != nil {
			l.Errorf("getting signature for %s: %s", filePath, err.Error())
		}

		hl = append(hl, mbtemplate.HeadLink{
			HRef:        pathStatic + path,
			Rel:         "stylesheet",
			CrossOrigin: "anonymous",
			Integrity:   signature,
		})
	}

	// generate head links
	var fs []mbtemplate.Script
	scriptPaths := []string{
		pathFileBootstrapJS,
	}
	for _, path := range scriptPaths {
		filePath := staticDir + path
		signature, err := getSignature(filePath)
		if err != nil {
			l.Errorf("getting signature for %s: %s", filePath, err.Error())
		}

		fs = append(fs, mbtemplate.Script{
			Src:         pathStatic + path,
			CrossOrigin: "anonymous",
			Integrity:   signature,
		})
	}

	return &Module{
		db:       db,
		language: lMod,
		minify:   m,
		store:    store,

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
func (m *Module) Route(s web.Server) error {
	staticFS, err := fs.Sub(megabot.Files, staticDir)
	if err != nil {
		return err
	}

	// Static Files
	s.PathPrefix(pathStatic + "/").Handler(http.StripPrefix(pathStatic+"/", http.FileServer(http.FS(staticFS))))

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
	protected.HandleFunc(pathChatbot, m.ChatbotGetHandler).Methods("GET")
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

	file, err := megabot.Files.Open(path)
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
