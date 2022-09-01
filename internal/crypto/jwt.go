package crypto

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type jwtClaims struct {
	Username           string `json:"username"`
	jwt.StandardClaims `json:"standardclaims"`
}

var jwtKey = []byte("pBNTRKr|a4<5xkn6x/,qu|+q)UT[F0=^")

// Creating jwt tokens
func GenerateJWT(username string) (string, error) {

	// Create claim
	claims := jwtClaims{Username: username}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

// Verifying and extracting information from jwt tokens
func ParseJWT(tokenString string) (bool, jwt.MapClaims, error) {

	// Get the token itself
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// If the signing method is wrong return an error
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	// If the token cannot be parsed
	if err != nil {
		return false, nil, err
	}

	// If claims cannot be extracted
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, claims, nil
	} else {
		return false, nil, errors.New("token invalid")
	}
}
