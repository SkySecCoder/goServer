package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"goServer/api"
	"goServer/auth"
	"goServer/banner"
	"goServer/base"
	"goServer/logs"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println(hello.Hello())
	/*
		Setting logging
	*/
	logFile, err := os.OpenFile("goServer.log", os.O_WRONLY|os.O_CREATE, 0755)
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

	/*
		Setting TLS along with supported cipher suites
	*/
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	/*
		Handling routes...
	*/
	router := mux.NewRouter()
	router.HandleFunc("/", base.Base)
	router.HandleFunc("/api", api.RequestHandler)
	router.HandleFunc("/api/{key}", api.RequestHandler)
	router.HandleFunc("/auth", auth.RequestHandler)
	router.HandleFunc("/logs", logs.RequestHandler)
	http.Handle("/", router)

	/*
		Setting server configurations:
		Addr:         Specifying port number and address
		Handler:      Specifying server handler for routes
		TLSConfig:    Specifying TLS configurations
		ReadTimeout:  Read timeout time
		WriteTimeout: Write timeout time
		IdleTimeout:  Idle Timeout time
	*/
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		TLSConfig:    cfg,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	if err := srv.ListenAndServeTLS("./tls/tls.crt", "./tls/tls.key"); err != nil {
		panic(err)
	}
}
