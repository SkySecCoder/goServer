package api

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"encoding/json"
)

type GenericResponse struct {
	Path 	string
	Data 	[]string
}

func RawGet(w http.ResponseWriter, r *http.Request) []byte {
	log.Info("[+] rawGet handling request...")
	var returnData GenericResponse

	returnData.Path = r.URL.Path
	returnData.Data = append(returnData.Data, "rawGet", "echoPost")

	data,_ := json.MarshalIndent(returnData, "", "    ")
	return data
}
