package main

import (
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
	"goServer/banner"
	"goServer/api"
	"goServer/base"
)

func main() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	fmt.Println(hello.Hello())
	http.HandleFunc("/", base.Base)
	http.HandleFunc("/api", rawPost.RawPost)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
