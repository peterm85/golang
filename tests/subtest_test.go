package tests

import "testing"

func Test_Sum_TestSuite(t *testing.T) {

	//setting up code

	t.Run("All positive inputs", func(t *testing.T) {
		if got := Sum(5, 8); got != 13 {
			t.Errorf("expected to get %d: got %d", 13, got)
		}
	})

	t.Run("Wrong use case", func(t *testing.T) {
		if got := Sum(2, 10); got != 15 {
			t.Errorf("expected to get %d: got %d", 15, got)
		}
	})

	//tear down code
}
