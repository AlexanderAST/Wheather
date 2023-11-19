package handler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() error {
	r := mux.NewRouter()

	r.HandleFunc("/hello", sayHello("АОАООАОАОАОАОАО")).Methods(http.MethodGet)
	r.HandleFunc("/wheather", takeWheather()).Methods(http.MethodPost)
	log.Println("Server started")

	return http.ListenAndServe(":8080", r)
}
