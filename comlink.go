package comlink

import (
	"net/http"
)

// Request is necessary data for http request
type Request struct {
	Client   *http.Client
	Method   string
	Path     string
	Payload  interface{}
	Response interface{}
}

type Dog struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// HTTPRequest provides http call
func HTTPRequest(req *Request) error {
	req.Response = Dog{
		Name: "Muttley",
		Age:  50,
	}

	return nil
}
