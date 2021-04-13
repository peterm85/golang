package tests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/peterm85/golang/tests/mocks"
)

func Test_GoMock(t *testing.T) {
	//settingUp
	const url string = "http://localhost:1234"

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockIClient := mocks.NewMockIClient(mockCtrl)
	// Expect Do to be called once with "http://localhost:1234" as parameter, and return http.Response from the mocked call.
	mockIClient.EXPECT().Get(url+"/some/path").Return(&http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`OK`))}, nil).Times(1)

	// given a Client & URL from our local test server
	api := API{
		Client:  mockIClient,
		BaseURL: url,
	}
	//when
	body, err := api.DoStuff()
	//then
	isOk(t, err)
	isEquals(t, []byte("OK"), body)
}
