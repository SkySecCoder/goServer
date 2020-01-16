package base

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"io/ioutil"
)

func Base(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Info("[+] "+r.Method+" Request at : "+path)
	message := "<h1>Placeholder goServer main page...</h1>"
	data, _ := ioutil.ReadAll("base.html")
	message = string(data)
	w.Write([]byte(message))
}
