package reverseproxy

import (
	"testing"
	"net/http"
)

func Test_AllowQueryTypesInspector_NotAllowedArgs_ReturnsError(t *testing.T) {
	allowedQueryTypes := []string{"access_token"}

	req, err := http.NewRequest("GET",
		"http://localhost:5000/test?q=counter&access_token=0123456789", nil)
	if err != nil {
		t.Errorf("Unable to create new http request.")
	}

	authinsp := NewAllowQueryTypesInspector(allowedQueryTypes)
	err = authinsp.Inspect(req)

	if err == nil {
		t.Errorf("Error should be returned since q is not an allowed query parameter.")
	}
}

func Test_AllowQueryTypesInspector_AllowedArg_NoErrorReturned(t *testing.T) {
	allowedQueryTypes := []string{"access_token"}

	req, err := http.NewRequest("GET",
		"http://localhost:5000/test?access_token=0123456789", nil)
	if err != nil {
		t.Errorf("Unable to create new http request.")
	}

	authinsp := NewAllowQueryTypesInspector(allowedQueryTypes)
	err = authinsp.Inspect(req)

	if err != nil {
		t.Errorf("Error should not be returned since access_token is allowed as a query parameter.")
	}
}

func Test_AllowQueryTypesInspector_AllowedArgs_NoErrorReturned(t *testing.T) {
	allowedQueryTypes := []string{"access_token", "page"}

	req, err := http.NewRequest("GET",
		"http://localhost:5000/test?page=1&access_token=0123456789", nil)
	if err != nil {
		t.Errorf("Unable to create new http request.")
	}

	authinsp := NewAllowQueryTypesInspector(allowedQueryTypes)
	err = authinsp.Inspect(req)

	if err != nil {
		t.Errorf("Error should not be returned since both access_token and page is allowed as a query parameters.")
	}
}
