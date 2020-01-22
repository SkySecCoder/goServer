package logs

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"encoding/json"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Info("[+] "+r.Method+" Request at : "+path)
	var response []byte

	if r.Method == http.MethodGet {
		response = LogGet(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		message := map[string]string{"Error":"405 "+r.Method+" is not allowed on "+path+" ..."}
		response,_ = json.MarshalIndent(message, "", "    ")
	}

	w.Write(response)
}
