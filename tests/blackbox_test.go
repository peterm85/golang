package tests_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/peterm85/golang/tests"
)

func Test_BB_DoStuffWithTestServer(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Write([]byte(`OK`))
	}))
	fmt.Println("Server available on: ", server.URL)
	// Close the server when test finishes
	defer server.Close()

	// given a Client & URL from our local test server
	api := tests.API{
		Client:  server.Client(),
		BaseURL: server.URL,
	}
	//when
	body, err := api.DoStuff()
	//then
	tests.IsOk(t, err)
	tests.IsEquals(t, []byte("OK"), body)
}

/*func Test_internalFunc(t *testing.T) {
	//when
	spec := internalFunc()
	//then
	IsEquals(t, 4, spec)
}*/
