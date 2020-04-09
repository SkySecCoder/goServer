package api

import (
	"encoding/json"
	"net/http"
)

type GenericResponse struct {
	Path string
	Data []string
}

func RawGet(w http.ResponseWriter, r *http.Request) []byte {
	var returnData GenericResponse

	returnData.Path = r.URL.Path
	returnData.Data = append(returnData.Data, "needs to be worked on to echo headers?")

	data, _ := json.MarshalIndent(returnData, "", "    ")
	return data
}
