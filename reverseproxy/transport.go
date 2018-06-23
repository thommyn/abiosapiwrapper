package reverseproxy

import (
	"net/http"
)

type Transporter interface {
	Get() http.RoundTripper
}

type HttpTransport struct {
	transport http.RoundTripper
}

func NewHttpTransport() Transporter {
	return &HttpTransport{
		transport: http.DefaultTransport,
	}
}

func (ht HttpTransport) Get() http.RoundTripper {
	return ht.transport
}
