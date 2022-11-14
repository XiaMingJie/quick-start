package main

import (
	"crypto/tls"
	"fmt"
	"github.com/lucas-clemente/quic-go/http3"
	"log"
	"net/http"
)

func loadTlsConf() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair("./cert/8804910_www.quickstart.top.pem", "./cert/8804910_www.quickstart.top.key")
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
	}, nil
}

func http3Handle() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Fprintf(writer, "welcome to quick start")
	})
	return mux
}

func main() {
	tlsConf, err := loadTlsConf()
	if err != nil {
		panic(err)
	}
	_ = tlsConf

	srv := &http3.Server{
		Addr: ":2333",
		//TLSConfig:      tlsConf,
		Handler:        http3Handle(),
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(srv.ListenAndServeTLS("./cert/8804910_www.quickstart.top.pem", "./cert/8804910_www.quickstart.top.key"))
}
