package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"github.com/SkySecCoder/goServer/goServer/banner"
)

func main() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	fmt.Println(hello.Hello())
	http.HandleFunc("/", echoHandler)
	http.HandleFunc("/api", apiHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
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

func apiHandler(w http.ResponseWriter, r *http.Request) {
	message := ""
	path := r.URL.Path
	log.Info("[+] Request at : "+path)

	if r.Method == http.MethodGet {
		message = "[-] Sorry only accept POST request on /api ...\n"
	} else {
		body,_ := ioutil.ReadAll(r.Body)
		message = string(body) + "\n"
	}

	w.Write([]byte(message))
}
