package reverseproxy

import (
	"net/http"
	"net/http/httputil"
	"../tokenbucket"
	"log"
)

type Handler interface {
	Get() func(http.ResponseWriter, *http.Request)
}

type httpHandler struct {
	p       *httputil.ReverseProxy
	reqinsp RequestInspector
	tb      tokenbucket.TokenBucket
}

func NewHttpHandler(p *httputil.ReverseProxy, reqinsp RequestInspector, tb tokenbucket.TokenBucket) Handler {
	return &httpHandler{
		p:       p,
		reqinsp: reqinsp,
		tb:      tb,
	}
}

func (hh httpHandler) Get() func(http.ResponseWriter, *http.Request) {
	return hh.handlerFunc
}

func (hh httpHandler) handlerFunc(w http.ResponseWriter, req *http.Request) {
	log.Println("request:", req.RemoteAddr, "want", req.RequestURI)

	// inspect request
	if err := hh.reqinsp.Inspect(req); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), HttpClientErrorForbidden)
		return
	}

	// consume one token and check validity
	// TODO: Rollback consume if error occurs in ServeHTTP?
	if err := hh.tb.ConsumeOneToken(); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), HttpClientErrorForbidden)
		return
	}

	// pass request...
	hh.p.ServeHTTP(w, req)
}
