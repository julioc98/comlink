package comlink

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	mockRequest Request
	client      *http.Client
	method      string
	path        string
	payload     Person
	// response     Dog
	responseMock Dog
)

func init() {
	mockRequest = Request{
		Client: &http.Client{},
	}
	responseMock = Dog{
		Name: "Muttley",
		Age:  50,
	}
}

func mockingServerCL() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case "GET":
			log.Println("Method: ", r.Method)
			resp, _ := json.Marshal(responseMock)
			fmt.Fprintln(w, string(resp))
		}
	}))

}

func Test_HTTPRequest_WithValidRequestMethodGET_ChangeCorrectResponse(t *testing.T) {

	mockRequest.Path = mockingServerCL().URL
	mockRequest.Method = "GET"

	_ = HTTPRequest(&mockRequest)
	respMock, _ := json.Marshal(responseMock)
	res, _ := json.Marshal(mockRequest.Response)

	if string(respMock) != string(res) {
		t.Errorf(" Response Mock: %s != Response HTTPRequest: %s", respMock, res)
	}
}
