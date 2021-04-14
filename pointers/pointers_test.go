package pointers

import (
	"testing"
)

func Test_Pointers(t *testing.T) {

	value := getInt()
	t.Log(value)

	valueP := &value
	t.Log(valueP)

	if value != *valueP {
		t.Errorf("expected to get %d: got %d", *valueP, value)
	}
}
