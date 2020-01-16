package main

import (
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/SkySecCoder/goServer/goServer/banner"
	"github.com/SkySecCoder/goServer/goServer/api"
	"github.com/SkySecCoder/goServer/goServer/base"
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
