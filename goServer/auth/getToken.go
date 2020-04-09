package auth

import (
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	log "github.com/sirupsen/logrus"
)

func GetToken(w http.ResponseWriter, r *http.Request) []byte {
	log.Info("[+] " + r.Method + " Request from " + r.RemoteAddr + " at : " + r.URL.Path)

	var mySigningKey = []byte("#REDACTED")

	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS512)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = true
	claims["name"] = "AakashPrasad"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Error("[-] " + r.Method + " Error " + r.RemoteAddr + " at : " + r.URL.Path)
		log.Error(err)
	}

	return []byte(tokenString)
}
