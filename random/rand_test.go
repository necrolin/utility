package random

import (
	"testing"
)

funct TestRand(t *testing.T {
	Init()
	
}

func TestRoll(t *testing.T) {
	Init()
	tests := 100

	vals := make([]int, 0, tests)
	for i := 0; i < tests; i++ {
		vals = append(vals, Roll())
	}

	for _, val := range vals {
		if val < 1 || val > 6 {
			t.Errorf("Roll() returned %q, expecting a number between 1 and 6", val)
		}
	}
}
