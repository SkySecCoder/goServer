package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"goServer/banner"
	"goServer/api"
	"goServer/base"
	"goServer/logs"
)

func main() {
	logFile, err := os.OpenFile("goServer.log", os.O_WRONLY | os.O_CREATE, 0755)
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	if err != nil {
		panic(err)
	}
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetReportCaller(true)
	log.SetFormatter(formatter)
	log.SetOutput(multiWriter)

	fmt.Println(hello.Hello())

	router := mux.NewRouter()
	router.HandleFunc("/", base.Base)
	router.HandleFunc("/api", api.RequestHandler)
	router.HandleFunc("/api/{key}", api.RequestHandler)
	router.HandleFunc("/logs", logs.RequestHandler)
	http.Handle("/", router)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
