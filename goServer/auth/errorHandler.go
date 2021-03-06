package auth

import (
	"encoding/json"
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	var response []byte

	if r.Method == http.MethodGet {
		response = GetToken(r)
	} else if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		message := map[string]string{"Error": "405 " + r.Method + " is not allowed on " + r.URL.Path + " ..."}
		response, _ = json.MarshalIndent(message, "", "    ")
	}

	w.Write(response)
}
