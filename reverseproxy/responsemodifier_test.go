package reverseproxy

import (
	"testing"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"bytes"
	"strconv"
	"crypto/md5"
)

type staticJsonConverter struct {}
func (staticJsonConverter) Convert(injson []interface{}) (outjson []interface{}, err error) {
	testJsonStr := `[{"id": 1}, {"id": 2}]`
	var testJson []interface{}
	json.Unmarshal([]byte(testJsonStr), &testJson)
	return testJson, nil
}

func hash(obj *http.Response) [16]byte {
	bytes, _ := json.Marshal(obj)
	return md5.Sum(bytes)
}

func getFakeHttpResponse() *http.Response {
	bodyBytes := []byte(`{"data": [{"id": 1}]}`)
	body := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	contentLength := len(bodyBytes)

	resp := &http.Response{
		Body: body,
		ContentLength: int64(len(bodyBytes)),
		StatusCode: 200,
		Header: make(http.Header, 0),
	}
	resp.Header.Set("Content-Length", strconv.Itoa(contentLength))

	return resp
}

func Test_ModifyResponseFunc_StatusCode402_ReturnsUnmodifiedBody(t *testing.T) {
	responseModifier := NewJsonConvResponseModifier(&staticJsonConverter{})
	fakeHttpResponse := getFakeHttpResponse()
	fakeHttpResponse.StatusCode = 402

	hash1 := hash(fakeHttpResponse)

	err := responseModifier.Get()(fakeHttpResponse)
	if err != nil {
		t.Errorf("There was un error calling ModifyResponseFunc.")
	}

	hash2 := hash(fakeHttpResponse)

	if hash1 != hash2 {
		t.Errorf("Http response should be the same after modification when status code is not 200.")
	}
}

func Test_ModifyResponseFunc_NilConverter_ReturnsUnmodifiedBody(t *testing.T) {
	responseModifier := NewJsonConvResponseModifier(nil)
	fakeHttpResponse := getFakeHttpResponse()
	hash1 := hash(fakeHttpResponse)

	err := responseModifier.Get()(fakeHttpResponse)
	if err != nil {
		t.Errorf("There was un error calling ModifyResponseFunc.")
	}

	hash2 := hash(fakeHttpResponse)

	if hash1 != hash2 {
		t.Errorf("Http response should be the same when a nil converter is specified.")
	}
}

func Test_ModifyResponseFunc_ReturnsCorrectModifedResponse(t *testing.T) {
	expectedBody := `{"data":[{"id":1},{"id":2}]}`
	expectedContentLength := len(expectedBody)
	expectedStatusCode := 200
	expectedContentLengthHeaderValue := strconv.Itoa(expectedContentLength)

	responseModifier := NewJsonConvResponseModifier(&staticJsonConverter{})
	fakeHttpResponse := getFakeHttpResponse()

	err := responseModifier.Get()(fakeHttpResponse)
	if err != nil {
		t.Errorf("There was un error calling ModifyResponseFunc.")
	}

	bodyBytes, _ := ioutil.ReadAll(fakeHttpResponse.Body)
	body := string(bodyBytes)
	contentLengthHeaderValue := fakeHttpResponse.Header.Get("Content-Length")

	if body != expectedBody {
		t.Errorf("Body of http respone is incorrect, got '%s' expected '%s'",
			body, expectedBody)
	}
	if fakeHttpResponse.StatusCode != expectedStatusCode {
		t.Errorf("Status code of http respone is incorrect, got '%d' expected '%d'",
			fakeHttpResponse.StatusCode, expectedStatusCode)
	}
	if fakeHttpResponse.ContentLength != int64(expectedContentLength) {
		t.Errorf("Content length of http respone is incorrect, got '%d' expected '%d'",
			fakeHttpResponse.ContentLength, expectedContentLength)
	}
	if contentLengthHeaderValue != expectedContentLengthHeaderValue {
		t.Errorf("Content length in header of http respone is incorrect, got '%s' expected '%s'",
			contentLengthHeaderValue, expectedContentLengthHeaderValue)
	}
}
