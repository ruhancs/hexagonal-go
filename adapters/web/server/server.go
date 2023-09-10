package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/ruhancs/hexagonal-go/adapters/web/handler"
	"github.com/ruhancs/hexagonal-go/application"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Server() {
	//mux ajuda atrabalhar com as rotas é um roteador
	router := mux.NewRouter()
	//negroni funciona com middleware
	middleware := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(router, middleware, w.Service)
	http.Handle("/", router)

	server:= &http.Server {
		ReadHeaderTimeout: 10 * time.Second,//tempo para pegar o header
		WriteTimeout: 10 * time.Second,// tempo para escrever
		Addr: ":9000",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr, "log:", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err!= nil {
		log.Fatal(err)
	}
}