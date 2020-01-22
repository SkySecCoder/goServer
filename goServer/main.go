package main

import (
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
	"goServer/banner"
	"goServer/api"
	"goServer/base"
	"os"
	"io"
)

//var logs []string

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
	http.HandleFunc("/", base.Base)
	http.HandleFunc("/api", api.RequestHandler)
	//http.HandleFunc("/log", log.RequestHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
