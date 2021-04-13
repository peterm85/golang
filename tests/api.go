package tests

import (
	"fmt"
	"io/ioutil"
)

type API struct {
	//Client *http.Client
	Client  IClient
	BaseURL string
}

func (api *API) DoStuff() ([]byte, error) {
	resp, err := api.Client.Get(api.BaseURL + "/some/path")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	internalFunc()

	// handling error and doing stuff with body that needs to be unit tested
	return body, err
}

func internalFunc() int {
	fmt.Println("Writting an internal log")
	return 4
}
