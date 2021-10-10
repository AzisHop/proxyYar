package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"proxyYar/requests"
)

func DbHandler(req *http.Request, scheme string) {
	var reqModel requests.Request
	reqModel.Method = req.Method
	reqModel.URLhost = req.URL.Host
	reqModel.URLscheme = scheme
	header := make(map[string]string, 0)
	for k, v := range req.Header {
		header[k] = v[0]
	}
	reqModel.Header = header
	body,_ := ioutil.ReadAll(req.Body)
	bodyString := string(body)
	reqModel.Body = bodyString
	reqModel.ContentLength = int(req.ContentLength)
	reqModel.Host = req.Host
	reqModel.RemoteAddr = req.RemoteAddr
	reqModel.RequestURI = req.RequestURI
	err := reqModel.SaveRequest()
	if err != nil {
		log.Println(err)
	}
}