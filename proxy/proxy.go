package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"proxyYar/certificate"
	"proxyYar/db"
	"proxyYar/handlers"
)

var rootCertificate certificate.Cert

func handleRequests(writer http.ResponseWriter, req *http.Request) {
	log.Println(req.Method)
	if req.Method == "GET" {
		handlers.HttpHandler(writer, req)
	} else {
		handlers.HttpsHandler(writer, req)
	}
}


func main() {

	db.CreateDataBaseConnection("docker", "docker", "localhost", "docker", 20)
	db.InitDataBase()

	server := &http.Server{
		Handler:      http.HandlerFunc(handleRequests),
		Addr:         ":8080",
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	rootCertificate = certificate.GetRootCertificate()

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}