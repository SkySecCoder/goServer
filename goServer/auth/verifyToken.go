package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	//log "github.com/sirupsen/logrus"
	"strings"
	"fmt"
	"net/http"
)

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	var mySigningKey = []byte("#REDACTED")

	/*
		Fetching value of Authorization header
	*/
	bearToken := r.Header.Get("Authorization")
	tokenString := strings.ReplaceAll(bearToken, "\n", "")
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
