package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"proxyYar/db"
	"proxyYar/handlers"
)

func main() {
	db.CreateDataBaseConnection("docker", "docker", "localhost", "docker", 20)
	r := mux.NewRouter()
	handlers.RepeaterClient = &http.Client{}
	r.HandleFunc("/request/{id}", handlers.MakeRequest).Methods("GET")
	r.HandleFunc("/requests", handlers.GetLastRequests).Methods("GET")

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal(err)
		return
	}

}