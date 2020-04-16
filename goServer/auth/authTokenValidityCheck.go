package auth

import (
	"net/http"
	"fmt"
)

func AuthTokenValidityCheck(w http.ResponseWriter, r *http.Request) {
	token, err := VerifyToken(r)
	if err != nil {
		fmt.Println(err)
		fmt.Println("[-] Token is invalid...")
		return
	}

	if token.Valid == false {
		fmt.Println(err)
	} else {
		fmt.Println("[+] Token is valid...")
	}
}
