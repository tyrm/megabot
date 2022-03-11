package webapp

import (
	"context"
	"github.com/go-http-utils/etag"
	"github.com/gorilla/sessions"
	"github.com/tyrm/megabot/internal/models"
	"golang.org/x/text/language"
	"net/http"
	"time"
)

// ResponseWriterX is a ResponseWriter that keeps track of status and body size
type ResponseWriterX struct {
	http.ResponseWriter
	status     int
	bodyLength int
}

// Write to the response writer, also updating body length
func (r *ResponseWriterX) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	if err != nil {
		return 0, err
	}
	r.bodyLength += n
	return n, nil
}

// WriteHeader sets the status of the response
func (r *ResponseWriterX) WriteHeader(status int) {
	r.ResponseWriter.WriteHeader(status)
	r.status = status
	return
}

// Middleware runs on every http request
func (m *Module) Middleware(next http.Handler) http.Handler {
	return etag.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := logger.WithField("func", "Middleware")
		start := time.Now()

		wx := &ResponseWriterX{
			ResponseWriter: w,
			status:         200,
			bodyLength:     0,
		}

		// Init Session
		us, err := m.store.Get(r, "megabot")
		if err != nil {
			l.Infof("got %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), sessionKey, us)

		// Retrieve our user and type-assert it
		val := us.Values["user"]
		if user, ok := val.(models.User); ok {
			ctx = context.WithValue(ctx, userKey, &user)
		}

		// create localizer
		lang := r.FormValue("lang")
		accept := r.Header.Get("Accept-Language")
		localizer, err := m.language.NewLocalizer(lang, accept)
		if err != nil {
			l.Errorf("could get localizer: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx = context.WithValue(ctx, localizerKey, localizer)

		// set request language
		ctx = context.WithValue(ctx, languageKey, getPageLang(lang, accept, m.language.Language().String()))

		// Do Request
		next.ServeHTTP(wx, r.WithContext(ctx))

		l.Debugf("rendering %s took %d ms", r.URL.Path, time.Since(start).Milliseconds())
	}), false)
}

func getPageLang(query, header, defaultLang string) string {
	l := logger.WithField("func", "getPageLang")

	if query != "" {
		t, _, err := language.ParseAcceptLanguage(query)
		if err == nil {
			l.Debugf("query languages: %v", t)
			if len(t) > 0 {
				l.Debugf("returning language: %s", t[0].String())
				return t[0].String()
			}
		} else {
			l.Debugf("query '%s' did not contain a valid lanaugae: %s", query, err.Error())
		}
	}

	if header != "" {
		t, _, err := language.ParseAcceptLanguage(header)
		if err == nil {
			l.Debugf("header languages: %v", t)
			if len(t) > 0 {
				l.Debugf("returning language: %s", t[0].String())
				return t[0].String()
			}
		} else {
			l.Debugf("query '%s' did not contain a valid lanaugae: %s", query, err.Error())
		}
	}

	l.Debugf("returning default language: %s", defaultLang)
	return defaultLang
}

// MiddlewareRequireAuth will redirect a user to login page if user not in context
func (m *Module) MiddlewareRequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		us := r.Context().Value(sessionKey).(*sessions.Session)

		if r.Context().Value(userKey) == nil {
			// Save current page
			us.Values["login-redirect"] = r.URL.Path
			err := us.Save(r, w)
			if err != nil {
				m.returnErrorPage(w, r, http.StatusInternalServerError, err.Error())
				return
			}

			// redirect to login
			http.Redirect(w, r, pathBase+pathLogin, http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
