package errhandling_test

import (
	"testing"

	eh "github.com/peterm85/golang/errhandling"
)

func Test_Listen(t *testing.T) {
	_, err := eh.Listen("ocalhost", 8080) // Ups: wrong host

	if err != nil {
		t.Errorf("%v", err)
	}
}
