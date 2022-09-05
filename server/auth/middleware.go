package auth

import (
	"context"
	"net/http"

	"github.com/unrolled/render"
	"rei.io/rei/internal/crypto"
	"rei.io/rei/internal/database"
	"rei.io/rei/internal/helpers"
)

// Middleware to check for jwt and session authenticity
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// First check if tokenstring is present
		cookie, err := r.Cookie("jwt")
		if err != nil {
			render.New().JSON(w, 500, map[string]string{"Error": "token not found"})
			return
		}

		tokenString := cookie.Value

		// Then verify the token
		any, err := crypto.ParseJWT(tokenString)
		if err != nil {
			render.New().JSON(w, 500, map[string]string{"Error": "invalid token. Did it expire?"})
			return
		}

		// Unwrap context to obtain correct db connection string
		ctx := r.Context()

		// Assert type string
		connStr := ctx.Value(helpers.ConnectionString{}).(string)

		// Initialise db
		db := new(database.EntClient)
		db.Init("postgres", connStr)

		loggedIn, ip, err := db.QuerySession(any["username"].(string))

		/*
			Error, login state not found (technically only having one of these two is enough since return value
			causes both state, but left it so its more clear
		*/
		if err != nil || loggedIn == nil {
			render.New().JSON(w, 500, map[string]string{"Error": "something went wrong"})
			return

			// If logged in but current registered session on another ip
		} else if *loggedIn && *ip != r.RemoteAddr {

			// Clear cookie
			http.SetCookie(w, &http.Cookie{
				Name:  "jwt",
				Value: "",
				Path:  "/"})
			render.New().JSON(w, 500, map[string]string{"Error": "Other session active. Please login again"})
			return

			// Not logged in (very old expired cookie?)
		} else if !*loggedIn {

			// Clear cookie
			http.SetCookie(w, &http.Cookie{
				Name:  "jwt",
				Value: "",
				Path:  "/"})
			render.New().JSON(w, 500, map[string]string{"Error": "Try logging in again"})
			return

			// If logged in and on correct ip
		} else if *loggedIn && *ip == r.RemoteAddr {

			// Renew jwt expiration
			tokenString, err = crypto.GenerateJWT(any["username"].(string))
			if err != nil {
				render.New().JSON(w, 500, map[string]string{"Error": "Something went wrong"})
				return
			}

			// Set cookie on user
			http.SetCookie(w, &http.Cookie{
				Name:  "jwt",
				Path:  "/",
				Value: tokenString,
				//SameSite: http.SameSiteStrictMode,
				//HttpOnly: true,
				//Secure: true,
			})
		}

		// Finally pass the claim into our next handler
		ctx = context.WithValue(r.Context(), helpers.UsernameClaim{}, any["username"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Must add behind Authenticate
func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := ctx.Value(helpers.UsernameClaim{}).(string)

		// If the user is wrong banish them
		if user != "arthur" {
			render.New().JSON(w, 403, map[string]string{"Error": "admin only"})
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
