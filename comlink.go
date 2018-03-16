package comlink

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Request is necessary data for http request
type Request struct {
	Client   *http.Client
	Method   string
	Path     string
	Payload  interface{}
	Response *http.Response
}

// HTTPRequest provides http call
func HTTPRequest(req *Request) error {

	reqPayload, err := json.Marshal(req.Payload)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(req.Method, req.Path, strings.NewReader(string(reqPayload)))
	if err != nil {
		return err
	}

	client := req.Client
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	req.Response = response

	return nil
}
