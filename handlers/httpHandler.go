package handlers

import (
	"io"
	"log"
	"net/http"
	"proxyYar/utils"
)

func HttpHandler(writer http.ResponseWriter, req *http.Request) {
	DbHandler(req, "http")
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Printf("proxy error: %s, request: %+v", err.Error(), req)
		http.Error(writer, err.Error(), http.StatusServiceUnavailable)
		return
	}
	log.Printf("status code: %d", resp.StatusCode)
	defer resp.Body.Close()
	utils.CopyHeader(writer.Header(), resp.Header)
	writer.WriteHeader(resp.StatusCode)
	io.Copy(writer, resp.Body)
}
