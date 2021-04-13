package tests

import (
	"net/http"
	"testing"
)

func (api *API) DoStuffWithRetries(retries int) ([]byte, error) {
	var err error
	var body []byte

	for i := 0; i < retries; i++ {
		body, err = api.DoStuff()
	}
	// handling error and doing stuff with body that needs to be unit tested
	return body, err
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////

type entry struct {
	url string
	err error
}

type spyIClient struct {
	retries int
	entries []entry
	client  stubClient
}

func (p *spyIClient) Get(url string) (*http.Response, error) {
	r, er := p.client.Get(url)
	p.retries++
	p.entries = append(p.entries, entry{url, er})
	return r, er
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////

func Test_DoStuffWithStubServerWithRetries(t *testing.T) {
	//given
	const retries int = 10

	spy := spyIClient{client: stubClient{}}
	api := API{
		Client:  &spy,
		BaseURL: "localhost",
	}
	//when
	api.DoStuffWithRetries(retries)
	//then
	isEquals(t, retries, spy.retries)
	isEquals(t, retries, len(spy.entries))
}
