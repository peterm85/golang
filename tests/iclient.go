package tests

import "net/http"

type IClient interface {
	Get(url string) (*http.Response, error)
}
