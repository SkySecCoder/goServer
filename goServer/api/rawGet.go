package rawGet

import (
	"net/http"
)

func RawGet(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	message := ""
	if path == "/echo" {
		log.Info("[+] Request at : "+path)
		message = "[+] Hello from server...\n"
	} else {
		log.Info("[+] Request at : "+path)
		message = "[+] Hello, welcome to serve.go...\n"
		message += " |- a small server for testing scripts...\n"
		message += " |- more can be added here later...\n"
	}

	w.Write([]byte(message))
}
