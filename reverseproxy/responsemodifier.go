package reverseproxy

import (
	"net/http"
	"../jquery"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"strconv"
	"compress/gzip"
)

type ResponseModifier interface {
	Get() func(*http.Response) error
}

type jsonQueryResponseModifier struct {
	converter jquery.JsonQuery
}

func NewJsonConvResponseModifier(converter jquery.JsonQuery) ResponseModifier {
	return &jsonQueryResponseModifier {
		converter: converter,
	}
}

func (rm jsonQueryResponseModifier) Get() func(*http.Response) error {
	return rm.modifyResponseFunc
}

func (rm jsonQueryResponseModifier) modifyResponseFunc(resp *http.Response) error {
	// no converter, just return...
	if rm.converter == nil {
		return nil
	}

	// if response status code is not 200 OK, just return...
	if resp.StatusCode != 200 {
		return nil
	}

	newrespbody, err := rm.getConvertedResponseBody(resp)
	if err != nil {
		return err
	}

	// update body of response
	rm.updateResponseBody(resp, newrespbody)

	return nil
}

func (rm jsonQueryResponseModifier) getConvertedResponseBody(resp *http.Response) ([]byte, error) {
	// get json form body content
	injson, err := rm.readBodyJson(resp)
	if err != nil {
		return nil, err
	}

	// convert json with supplied jsonconv method
	data := injson["data"].([]interface{})
	outjson, err := rm.converter.GetSubNodes(data)
	if err != nil {
		return nil, err
	}

	// overwrite data node with new converted json
	injson["data"] = outjson
	newrespbody, err := json.Marshal(injson)
	if err != nil {
		return nil, err
	}

	return newrespbody, nil
}

func (rm jsonQueryResponseModifier) readBodyJson(resp *http.Response) (map[string]interface{}, error) {
	defer resp.Body.Close()

	decoder, err := rm.getBodyDecoder(resp)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (rm jsonQueryResponseModifier) getBodyDecoder(resp *http.Response) (*json.Decoder, error) {
	var decoder *json.Decoder

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer gz.Close()
		decoder = json.NewDecoder(gz)
	default:
		decoder = json.NewDecoder(resp.Body)
	}

	return decoder, nil
}

func (rm jsonQueryResponseModifier) updateResponseBody(resp *http.Response, newrespbody []byte) error {
	buf, err := rm.getBodyBytesBuffer(resp, newrespbody)
	if err != nil {
		return err
	}

	// set body content and body content length
	resp.Body = ioutil.NopCloser(buf)
	contentLength := buf.Len()

	// overwrite Content-Length if present in header
	resp.ContentLength = int64(contentLength)
	if resp.Header.Get("Content-Length") != "" {
		resp.Header.Set("Content-Length", strconv.Itoa(contentLength))
	}

	return nil
}

func (rm jsonQueryResponseModifier) getBodyBytesBuffer(resp *http.Response, newrespbody []byte) (buf *bytes.Buffer, err error) {
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		buf, err = rm.encodeContentAsGzip(newrespbody)
		if err != nil {
			return nil, err
		}
	default:
		buf = bytes.NewBuffer(newrespbody)
	}

	return buf, nil
}

func (rm jsonQueryResponseModifier) encodeContentAsGzip(data []byte) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(data); err != nil {
		return nil, err
	}

	if err := gz.Flush(); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}

	return &buf, nil
}
