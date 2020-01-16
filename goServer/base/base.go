package base

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Base(w http.RequestWriter, r *http.Request) {
	path := r.URL.Path
	log.Info("[+] "+r.Method+" Request at : "+path)
	message := "<h1>Placeholder goServer main page...</h1>"
	w.Write([]byte(message))
}
