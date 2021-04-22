package errhandling_test

import (
	"errors"
	"testing"

	eh "github.com/peterm85/golang/errhandling"
)

func Test_ListenAcc_Contain_Error(t *testing.T) {
	_, errs := eh.ListenAcc("ocalhost", 8080)

	target := errors.New("Listen: lookup ocalhost: no such host")
	if hasError(errs, target) {
		t.Errorf("%v", errs)
	}
}

func hasError(err []error, target error) bool {
	for _, e := range err {
		if errors.Is(e, target) {
			return true
		}
	}
	return false
}
