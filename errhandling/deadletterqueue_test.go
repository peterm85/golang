package errhandling_test

import (
	"testing"

	eh "github.com/peterm85/golang/errhandling"
)

func Test_Workspace(t *testing.T) {
	eh.RunWorkspace("Hello!")
}
