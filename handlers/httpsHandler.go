package handlers

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"proxyYar/certificate"
)

func HttpsHandler(writer http.ResponseWriter, req *http.Request) {
	log.Println(req)
	cert, err := certificate.CreateLeafCertificate(req.Host)
	if err != nil {
		log.Println(err)
	}
	tlsConfig := & tls.Config{
		Certificates:                []tls.Certificate{*cert},
		GetCertificate:              func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
			return certificate.CreateLeafCertificate(info.ServerName)
		},
	}
	destinationConnection, err := tls.Dial("tcp", req.Host, tlsConfig)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusServiceUnavailable)
		return
	}
	hijacker, ok := writer.(http.Hijacker)
	if !ok {
		http.Error(writer, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	clientConnection, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusServiceUnavailable)
	}
	_, err = clientConnection.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))

	tlsConnection := tls.Server(clientConnection, tlsConfig)
	err = tlsConnection.Handshake()

	go transfer(destinationConnection, tlsConnection, true, req.URL.Host)
	go transfer(tlsConnection, destinationConnection, false, req.URL.Host)
}

func copyToDB(buffer *bytes.Buffer, requestHost string) {
	p := make([]byte, 1024*1024*8)

	for {
		n, err := buffer.Read(p)
		if err != nil{
			if err == io.EOF {
				fmt.Println(string(p[:n])) //should handle any remaining bytes.
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(p[:n]))
	}
	reader:=bufio.NewReader(bytes.NewReader(p))
	req, err := http.ReadRequest(reader)
	if err != nil {
		log.Println(err)
		return
	} else {
		req.URL.Host = requestHost
		DbHandler(req, "https")
	}
}


func transfer(destination io.WriteCloser, source io.ReadCloser, copy bool, requestHost string) {
	defer destination.Close()
	defer source.Close()
	if copy {
		buffer := &bytes.Buffer{}
		duplicateSources := io.MultiWriter(destination, buffer) //we copy data from source into buffer and destination
		io.Copy(duplicateSources, source)
		go copyToDB(buffer,  requestHost)

	} else {
		io.Copy(destination, source)
	}
}
