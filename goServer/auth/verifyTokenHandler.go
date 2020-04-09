package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

func VerifyToken() {
	var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256
	})
}
