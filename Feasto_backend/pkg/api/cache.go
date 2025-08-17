package api

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
)

// Global cache: 5 min default expiration, cleanup every 10 min
var c = cache.New(5*time.Minute, 10*time.Minute)

// Custom responseWriter to capture response body and status
type responseWriter struct {
	http.ResponseWriter
	buf        bytes.Buffer
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	rw.buf.Write(data)
	return rw.ResponseWriter.Write(data)
}

// CacheMiddleware wraps GET requests and caches successful responses
func CacheMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Only cache GET and HEAD requests
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			next(w, r)
			return
		}

		// Generate cache key
		var cacheKey string
		if strings.HasPrefix(r.URL.Path, "/menu/cat/") {
			vars := mux.Vars(r)
			id := vars["id"]
			cacheKey = "/menu/cat/" + id
		} else {
			cacheKey = r.URL.String()
		}

		// Check cache
		if cachedResponse, found := c.Get(cacheKey); found {
			w.Header().Set("Content-Type", "application/json")
			w.Write(cachedResponse.([]byte))
			return
		}

		// Capture response
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next(rw, r)

		// Only cache successful responses (2xx)
		if rw.statusCode >= 200 && rw.statusCode < 300 {
			c.Set(cacheKey, rw.buf.Bytes(), cache.DefaultExpiration)
		}
	}
}
