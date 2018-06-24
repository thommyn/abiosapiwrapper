package reverseproxy

import (
	"net/http"
	"net/url"
	"fmt"
)

type Director interface {
	Get() func(*http.Request)
}

type targetDirector struct {
	target *url.URL
}

func NewTargetDirector(target *url.URL) Director {
	return &targetDirector{
		target: target,
	}
}

func (td targetDirector) Get() func(*http.Request) {
	return td.directorFunc
}
func (td targetDirector) directorFunc(req *http.Request) {
	req.URL.Host = td.target.Host
	req.URL.Scheme = td.target.Scheme
	req.URL.Path = td.target.Path

	// combine request and target raw queries
	if td.target.RawQuery != "" {
		req.URL.RawQuery += fmt.Sprintf("&%s", td.target.RawQuery)
	}
}
