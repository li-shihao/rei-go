package auth

import (
	"context"
	"net/http"

	"github.com/unrolled/render"
	"rei.io/rei/internal/crypto"
)

// Middleware to check for jwt authenticity
func JWTAuthenticate(next http.Handler) http.Handler {

	type UsernameJWT struct{}

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
		// Finally pass the claim into our next handler
		ctx := context.WithValue(r.Context(), UsernameJWT{}, any["username"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
