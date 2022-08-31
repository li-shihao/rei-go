package crypto

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("pBNTRKr|a4<5xkn6x/,qu|+q)UT[F0=^")

func GenerateJWT(username string) (string, time.Time, error) {
	type jwtClaims struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := jwtClaims{Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, expirationTime, err
}
