package api

import (
	"net/http"
	"encoding/json"
)

func HelpGet(w http.ResponseWriter, r *http.Request) []byte {
	var returnData GenericResponse

	returnData.Path = r.URL.Path
	returnData.Data = append(returnData.Data, "rawGet", "echoPost", "helpGet")

	data,_ := json.MarshalIndent(returnData, "", "    ")
	return data
}
