/*
Package auth generates the JWT Token for the client to be authenticated when rquesting data from API
*/
package auth

import (
	"PGL/Client/log"
	"os"

	"github.com/golang-jwt/jwt"
)

//Generates the JWT token for the token header
func GenerateJWT() (string, error) {

	key := os.Getenv("JWT_SECRETKEY")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Client"

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		log.Warning.Println(err)
		return "", err
	}

	return tokenString, nil
}
