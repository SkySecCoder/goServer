package logs

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

func LogGet(w http.ResponseWriter, r *http.Request) []byte {
	var data []byte

	file, err := os.Open("goServer.log")
	if err != nil {
		data = []byte(err.Error())
		log.Error("[-] " + err.Error())
		return data
	}
	data, _ = ioutil.ReadAll(file)
	file.Close()

	return data
}
