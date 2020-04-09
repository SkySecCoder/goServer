package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Path string
	Data map[string]interface{}
}

type ResponseError struct {
	Path string
	Data string
}

func EchoPost(w http.ResponseWriter, r *http.Request) []byte {
	var returnData Response
	var data []byte

	body, _ := ioutil.ReadAll(r.Body)
	var checker map[string]interface{}
	if json.Unmarshal([]byte(body), &checker) == nil {
		var dat map[string]interface{}
		json.Unmarshal(body, &dat)
		returnData.Data = dat
		returnData.Path = r.URL.Path

		data, _ = json.MarshalIndent(returnData, "", "    ")
	} else {
		message := "This is not a json..."
		var errorData ResponseError
		errorData.Data = message
		errorData.Path = r.URL.Path

		data, _ = json.MarshalIndent(errorData, "", "    ")
	}

	return data
}
