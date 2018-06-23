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

type HttpHandler struct {
	p *httputil.ReverseProxy
	reqinsp RequestInspector
	tb tokenbucket.TokenBucket
}

func NewHttpHandler(p *httputil.ReverseProxy, reqinsp RequestInspector,
	tb tokenbucket.TokenBucket) Handler {
	return &HttpHandler{
		p: p,
		reqinsp: reqinsp,
		tb: tb,
	}
}

func (hh HttpHandler) Get() func(http.ResponseWriter, *http.Request) {
	return hh.handlerFunc
}

func (hh HttpHandler) handlerFunc(w http.ResponseWriter, req *http.Request) {
	log.Println("request:", req.RemoteAddr, "want", req.RequestURI)

	// inspect request
	if err := hh.reqinsp.Inspect(req); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 403)
		return
	}

	// consume one token and check validity
	// TODO: Rollback consume if error occurs in ServeHTTP?
	if err := hh.tb.ConsumeOneToken(); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 403)
		return
	}

	// pass request...
	hh.p.ServeHTTP(w, req)
}
