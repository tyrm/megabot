package webapp

import (
	"net/http"
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wx := &ResponseWriterX{
			ResponseWriter: w,
			status:         200,
			bodyLength:     0,
		}

		// Do Request
		next.ServeHTTP(wx, r)
	})
}
