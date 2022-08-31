package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	api "rei.io/rei/server/API"
	"rei.io/rei/server/auth"
)

// Holder for db connection string
var connStr string

func CreateServer(str string) *chi.Mux {

	// Set db connection string from parameter
	connStr = str

	csrfMiddleware := csrf.Protect([]byte("ir0LFQIIHiWbwGZlbkAqFGPcCGJi0U8k"))

	r := chi.NewRouter()

	/*
		pprof profiling endpoint
		To use, go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=60
		Flags for heap include inuse_space, inuse_objects, alloc_space, alloc_objects
	*/
	r.Mount("/debug", middleware.Profiler())

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(setDB)
		r.Use(csrfMiddleware)
		r.Get("/txcount", api.TotalTransactionCount)
	})

	r.Route("/signup", func(r chi.Router) {
		r.Use(setDB)
		r.Post("/", auth.Signup)
	})

	r.Route("/login", func(r chi.Router) {
		r.Use(setDB)
		r.Post("/", auth.Login)
	})

	return r
}

// Middleware to pack correct db connection string to our handlers
func setDB(next http.Handler) http.Handler {

	type ConnectionString struct{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ConnectionString{}, connStr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
