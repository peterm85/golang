package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_DoStuffWithTestServer(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Write([]byte(`OK`))
	}))
	fmt.Println("Server available on: ", server.URL)
	// Close the server when test finishes
	defer server.Close()

	// given a Client & URL from our local test server
	api := API{
		Client:  server.Client(),
		BaseURL: server.URL,
	}
	//when
	body, err := api.DoStuff()
	//then
	isOk(t, err)
	isEquals(t, []byte("OK"), body)
}

func Test_internalFunc(t *testing.T) {
	//when
	spec := internalFunc()
	//then
	isEquals(t, 4, spec)
}
