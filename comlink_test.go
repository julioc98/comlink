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

type Dog struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	mockRequest  Request
	client       *http.Client
	method       string
	path         string
	payload      Person
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
	payload = Person{
		Name: "Richard Milhous Dastardly",
		Age:  50,
	}
}

func mockingServerCL() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case "GET":
			log.Println("Method: ", r.Method)
			resp, _ := json.Marshal(responseMock)
			log.Println("Resp: ", string(resp))
			fmt.Fprintln(w, string(resp))

		case "POST":
			log.Println("Method: ", r.Method)

			var body Person

			defer r.Body.Close()
			json.NewDecoder(r.Body).Decode(&body)

			person, _ := json.Marshal(body)
			log.Println("Payload: ", string(person))

			resp, _ := json.Marshal(responseMock)
			log.Println("Resp: ", string(resp))
			fmt.Fprintln(w, string(resp))
		}
	}))

}

func Test_HTTPRequest_WithValidRequestMethodGET_ChangeCorrectResponse(t *testing.T) {
	var response Dog
	mockRequest.Path = mockingServerCL().URL
	mockRequest.Method = "GET"

	_ = HTTPRequest(&mockRequest)

	defer mockRequest.Response.Body.Close()
	json.NewDecoder(mockRequest.Response.Body).Decode(&response)
	resp, _ := json.Marshal(response)

	respMock, _ := json.Marshal(responseMock)

	if string(respMock) != string(resp) {
		t.Errorf(" Response Mock: %s != Response HTTPRequest: %s", respMock, resp)
	}
}

func Test_HTTPRequest_WithValidRequestMethodPOST_ChangeCorrectResponse(t *testing.T) {
	var response Dog
	mockRequest.Path = mockingServerCL().URL
	mockRequest.Method = "POST"
	mockRequest.Payload = payload

	_ = HTTPRequest(&mockRequest)

	defer mockRequest.Response.Body.Close()
	json.NewDecoder(mockRequest.Response.Body).Decode(&response)
	resp, _ := json.Marshal(response)

	respMock, _ := json.Marshal(responseMock)

	if string(respMock) != string(resp) {
		t.Errorf(" Response Mock: %s != Response HTTPRequest: %s", respMock, resp)
	}
}
