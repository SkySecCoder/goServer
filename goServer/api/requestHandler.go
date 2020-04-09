package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Info("[+] " + r.Method + " Request from " + r.RemoteAddr + " at : " + path)
	var response []byte

	requestVars := mux.Vars(r)
	if (requestVars["key"] == "echoPost") && (r.Method == http.MethodPost) {
		response = EchoPost(w, r)
	} else if (requestVars["key"] == "rawGet") && (r.Method == http.MethodGet) {
		response = RawGet(w, r)
	} else if (r.Method == http.MethodGet) && (strings.ContainsAny(path[1:len(path)], "/") == false) {
		response = HelpGet(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		message := map[string]string{"Error": "405 " + r.Method + " is not allowed on " + path + " ..."}
		response, _ = json.MarshalIndent(message, "", "    ")
	}

	w.Write(response)
}
