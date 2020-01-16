package base

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"io/ioutil"
	"os"
)

func Base(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Info("[+] "+r.Method+" Request at : "+path)
	message := "<h1>Placeholder goServer main page...</h1>"

	file, _ := os.Open("base.html")
	data, _ := ioutil.ReadAll(file)

	message = string(data)
	w.Write([]byte(message))
}
