package tests

import "testing"

func Test_Sum_Hybrid(t *testing.T) {

	//Case definition
	specs := []struct {
		descr  string
		inputX int
		inputY int
		exp    int
	}{
		{descr: "All positive inputs", inputX: 5, inputY: 8, exp: 13},
		{descr: "Positive and negative inputs", inputX: 5, inputY: -3, exp: 2},
		{descr: "Negative results", inputX: 5, inputY: -12, exp: -7},
		{descr: "All negative inputs", inputX: -5, inputY: -12, exp: -17},
		//{descr: "Wrong use case", inputX: 2, inputY: 10, exp: 15},
	}

	//Test-runner code
	for spectIndex, spec := range specs {
		t.Run(spec.descr, func(t *testing.T) {
			//t.Parallel()
			//setting up code
			if got := Sum(spec.inputX, spec.inputY); got != spec.exp {
				t.Errorf("[spec %d: %s] expected to get %d: got %d", spectIndex, spec.descr, spec.exp, got)
			}
			//tearring down code
		})
	}
}
