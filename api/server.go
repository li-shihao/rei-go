package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type key struct{}

// Holder for db connection string
var connStr string

func CreateServer(str string) *chi.Mux {

	// Set db connection string from parameter
	connStr = str

	r := chi.NewRouter()

	/*
		pprof profiling endpoint
		To use, go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=60
		Flags for heap include inuse_space, inuse_objects, alloc_space, alloc_objects
	*/
	r.Mount("/debug", middleware.Profiler())

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(setDB)
		r.Get("/txcount", TotalTransactionCount)
	})

	return r
}

// Middleware to pack correct db connection string to our handlers
func setDB(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), key{}, connStr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
