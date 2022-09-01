package auth

import (
	"context"
	"net/http"

	"github.com/unrolled/render"
	"rei.io/rei/internal/crypto"
	"rei.io/rei/internal/database"
)

type UsernameJWT struct{}

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
		token, any, err := crypto.ParseJWT(tokenString)
		if err != nil || !token {
			render.New().JSON(w, 500, map[string]string{"Error": "bad token"})
			return
		}

		// Unwrap context to obtain correct db connection string
		ctx := r.Context()

		// Assert type string
		connStr := ctx.Value(ConnectionString{}).(string)

		// Initialise db
		db := new(database.EntClient)
		db.Init("postgres", connStr)

		loggedIn, ip, err := db.QuerySession(any["username"].(string))
		if err != nil || loggedIn == nil {
			render.New().JSON(w, 500, map[string]string{"Error": "something went wrong"})
			return
		} else if *loggedIn && *ip != r.RemoteAddr {
			http.SetCookie(w, &http.Cookie{
				Name:  "jwt",
				Value: ""})
			render.New().JSON(w, 500, map[string]string{"Error": "logged off from another client. Please login again"})
			return
		} else if !*loggedIn {
			http.SetCookie(w, &http.Cookie{
				Name:  "jwt",
				Value: ""})

			render.New().JSON(w, 500, map[string]string{"Error": "Try logging in again"})
			return
		}

		// Finally pass the claim into our next handler
		ctx = context.WithValue(r.Context(), UsernameJWT{}, any["username"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
