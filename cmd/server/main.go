package main

import (
	"fmt"
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
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "welcome to quick start")
	})
	//log.Fatal(http3.ListenAndServeQUIC(":2333", "cert/8804910_www.quickstart.top.pem", "cert/8804910_www.quickstart.top.key", nil))
	log.Fatalln(http.ListenAndServeTLS(":2333", "cert/8804910_www.quickstart.top.pem", "cert/8804910_www.quickstart.top.key", nil))
}
