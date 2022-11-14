package main

import (
	"fmt"
	"github.com/lucas-clemente/quic-go/http3"
	"log"
	"net/http"
)

func http3Handle() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "welcome to quick start")
	})
	return mux
}

func main() {
	srv := &http3.Server{
		Addr:           ":2333",
		Handler:        http3Handle(),
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(srv.ListenAndServeTLS("./cert/8804910_www.quickstart.top.pem", "./cert/8804910_www.quickstart.top.key"))
}
