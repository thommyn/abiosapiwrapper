package reverseproxy

import (
	"testing"
	"net/url"
	"net/http"
)

func Test_DirectorFunc_NotAllowedArgs_ReturnsError(t *testing.T) {
	var expectedRawQuery string = "q1=1&q2=2&q3=3&q4=4"

	target := &url.URL{
		Host: "api.target.com",
		Scheme: "https",
		Path: "/apitarget/testtarget",
		RawQuery: "q3=3&q4=4",
	}

	request := &http.Request{
		Host: "",
		URL: &url.URL {
			Host: "api.request.com",
			Scheme: "https",
			Path: "/apireq/testreq",
			RawQuery: "q1=1&q2=2",
		},
	}

	director := NewTargetDirector(target)
	directorFunc := director.Get()
	directorFunc(request)

	if request.URL.Host != target.Host {
		t.Errorf("Request host does not match target, got %s, expected %s", request.URL.Host, target.Host)
	}
	if request.URL.Scheme != target.Scheme {
		t.Errorf("Request scheme does not match target, got %s, expected %s", request.URL.Scheme, target.Scheme)
	}
	if request.URL.Path != target.Path {
		t.Errorf("Request path does not match target, got %s, expected %s", request.URL.Path, target.Path)
	}
	if request.URL.RawQuery != expectedRawQuery {
		t.Errorf("Request raw query does not match target, got %s, expected %s", request.URL.RawQuery, expectedRawQuery)
	}
}
