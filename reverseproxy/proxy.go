package reverseproxy

import (
	"net/http/httputil"
)

func NewReverseProxy(director Director, responseModifier ResponseModifier, transporter Transporter) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: director.Get(),
		ModifyResponse: responseModifier.Get(),
		Transport: transporter.Get(),
	}
}
