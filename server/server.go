package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/csrf"
	"rei.io/rei/internal/helpers"
	"rei.io/rei/server/api"
	"rei.io/rei/server/auth"
)

// Holder for db connection string

var connStr string

func CreateServer(str string) *chi.Mux {

	// Set db connection string from parameter
	connStr = str

	csrfMiddleware := csrf.Protect([]byte("ir0LFQIIHiWbwGZlbkAqFGPcCGJi0U8k"))

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	/*
		pprof profiling endpoint
		To use, go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=60
		Flags for heap include inuse_space, inuse_objects, alloc_space, alloc_objects
	*/
	r.Mount("/debug", middleware.Profiler())

	r.Route("/", func(r chi.Router) {
		r.Use(setDB)
		r.Post("/signup", auth.Signup)
		r.Post("/login", auth.Login)
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(setDB)
		r.Use(csrfMiddleware)
		r.Use(auth.Authenticate)
		r.Get("/txcount", api.TotalTransactionCount)
		r.Post("/mock", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	})

	r.Route("/admin", func(r chi.Router) {
		r.Use(setDB)
		r.Use(auth.AdminOnly)
	})

	return r
}

// Middleware to pack correct db connection string to our handlers
func setDB(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), helpers.ConnectionString{}, connStr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
