package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Db     *gorm.DB
	Router *mux.Router
}

func (server *Server) InitalizeRouters() {
	server.Router.HandleFunc("/", server.Home).Methods("GET")
	server.Router.HandleFunc("/products", server.Product).Methods("GET")

	staticFile := http.Dir("./assets")
	statisHandler := http.StripPrefix("/public", http.FileServer(staticFile))
	server.Router.PathPrefix("/public").Handler(statisHandler).Methods("GET")
}
