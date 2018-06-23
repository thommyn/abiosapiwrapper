package reverseproxy

import (
	"net/http/httputil"
	"net/http"
)

func NewReverseProxy(director Director, responseModifier ResponseModifier, transporter Transporter) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: director.Get(),
		ModifyResponse: responseModifier.Get(),
		Transport: transporter.Get(),
	}
}

func Handle(handler Handler) func(http.ResponseWriter, *http.Request) {
	return handler.Get()
}
