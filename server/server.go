package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"rei.io/rei/internal/helpers"
	"rei.io/rei/server/api"
	"rei.io/rei/server/auth"
)

// Holder for db connection string

var connStr string

func CreateServer(str string) *chi.Mux {

	// Set db connection string from parameter
	connStr = str

	// TODO: CSRF
	//csrfMiddleware := csrf.Protect([]byte("ir0LFQIIHiWbwGZlbkAqFGPcCGJi0U8k"))

	r := chi.NewRouter()

	// CORS (I need this for anything to work)
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:3000"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Origin"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	/*
		pprof profiling endpoint
		To use, go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=60
		Flags for heap include inuse_space, inuse_objects, alloc_space, alloc_objects
	*/
	r.Mount("/debug", middleware.Profiler())

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Use(setDB)
		r.Post("/signup", auth.Signup)
		r.Post("/login", auth.Login)
		r.Handle("/query", api.GraphQLHandler(connStr))
		r.Handle("/playground", api.PlaygroundQLHandler("/query"))
	})

	// Private Routes (Requires Auth)
	r.Group(func(r chi.Router) {
		r.Use(setDB)
		r.Use(auth.Authenticate)
		r.Route("/api/v1", func(r chi.Router) {
			r.Get("/txcount", api.TotalTransactionCount)
		})
		r.Route("/admin", func(r chi.Router) {
			r.Use(auth.AdminOnly)
		})
		r.Post("/logout", auth.Logout)
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
