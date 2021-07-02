/*
Package auth provides the authentication middleware for the mux router
*/
package auth

import (
	"PGL/APIServer/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

//middleware to authenticate the client with JWT
func AuthJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := os.Getenv("JWT_SECRETKEY")
		tokenString := r.Header.Get("Token")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			//[]byte containing secret, e.g. []byte("my_secret_key")
			return []byte(key), nil
		})
		var res models.OtherRes
		if err != nil {
			res.Msg = "unexpected signing method"
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(res)
			return
		}
		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			res.Msg = "Not Authorized"
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(res)
		}
	})
}
