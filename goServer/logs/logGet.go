package logs

import (
	"os"
	"io/ioutil"
	"net/http"
	log "github.com/sirupsen/logrus"
	//"encoding/json"
)

func LogGet(w http.ResponseWriter, r *http.Request) []byte {
	var data []byte

	file, err := os.Open("goServer.log")
	if err != nil {
		log.Error("[-] Error opening file goServer.log")
	}
	data,_ = ioutil.ReadAll(file)
	file.Close()

	return data
}
