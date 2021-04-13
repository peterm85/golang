package tests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type stubClient struct{}

func (s *stubClient) Get(url string) (*http.Response, error) {
	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
	}, nil
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////

func Test_DoStuffWithStubServer(t *testing.T) {
	//given
	api := API{
		Client:  &stubClient{},
		BaseURL: "localhost",
	}
	//when
	body, err := api.DoStuff()
	//then
	isOk(t, err)
	isEquals(t, []byte("OK"), body)
}
