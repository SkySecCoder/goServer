package rawPost

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
)

type Response struct {
	Path 	string
	Data 	map[string]interface{}
}

type ResponseError struct {
	Path 	string
	Data 	string
}

func RawPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Info("[+] "+r.Method+" Request at : "+path)
	var returnData Response
	var data []byte

	if r.Method == http.MethodGet {
		message := "Sorry only accept POST request on /api ..."
		var errorData ResponseError
		errorData.Data = message
		errorData.Path = r.URL.Path
		
		data,_ = json.MarshalIndent(errorData, "", "    ")
	} else {
		body,_ := ioutil.ReadAll(r.Body)
		var checker map[string]interface{}
		if json.Unmarshal([]byte(body), &checker) == nil {
			var dat map[string]interface{}
			json.Unmarshal(body, &dat)
			returnData.Data = dat
			returnData.Path = r.URL.Path

			data,_ = json.MarshalIndent(returnData, "", "    ")
		} else {
			message := "This is not a json..."
			var errorData ResponseError
			errorData.Data = message
			errorData.Path = r.URL.Path
			
			data,_ = json.MarshalIndent(errorData, "", "    ")
		}
	}
	w.Write(data)
}
