package rawPost

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func RawPost(w http.ResponseWriter, r *http.Request) {
	message := ""
	path := r.URL.Path
	log.Info("[+] "+r.Method+" Request at : "+path)

	if r.Method == http.MethodGet {
		message = "[-] Sorry only accept POST request on /api ...\n"
	} else {
		body,_ := ioutil.ReadAll(r.Body)
		message = string(body) + "\n"
	}

	w.Write([]byte(message))
}
