package auth

import (
	"encoding/json"
	"net/http"
	//"fmt"
	"io/ioutil"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	database := map[string]string {
		"name": "aakash",
		"password": "qwerty",
	}

	response := []byte{}

	jsonData := map[string]string{}
	if r.Method != http.MethodPost {
		response = []byte(r.Method + " not allowed...")
		w.Write(response)
		return
	}
	data, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(data, &jsonData)
	if (database["name"] == jsonData["name"]) && (database["password"] == jsonData["password"]) {
		response = GetToken(r)
	} else {
		response = []byte("[-] Invalid credentials, auth failed...")
		w.Write(response)
		return
	}
	w.Write(response)
}
