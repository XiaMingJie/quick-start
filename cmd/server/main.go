package main

import (
	"github.com/lucas-clemente/quic-go/http3"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func http3Handle() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/hello", websocket.Handler(wsHello))
	return mux
}

func wsHello(conn *websocket.Conn) {
	conn.Write([]byte("hello! success"))
}

func main() {
	srv := http3.Server{
		Addr:    ":2333",
		Handler: http3Handle(),
	}

	log.Fatal(srv.ListenAndServeTLS("cert/8804910_www.quickstart.top.pem", "cert/8804910_www.quickstart.top.key"))
}
