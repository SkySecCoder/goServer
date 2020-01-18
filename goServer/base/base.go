package base

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"
	"strings"
)

func Base(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Info("[+] "+r.Method+" Request at : "+path)
	message := ""
	if strings.ContainsAny(path[1:len(path)], "/") {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"Error":"404 Page not found..."}
		data,_ := json.MarshalIndent(response, "", "    ")
		message = string(data)
	} else {
		file, _ := os.Open("base.html")
		data, _ := ioutil.ReadAll(file)
		message = string(data)
	}

	w.Write([]byte(message))
}
