package letters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounterLetter(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "Goooooooooddddd",
			Expected: "G1o9d5",
		},
		{
			Input:    "Yes",
			Expected: "Y1e1s1",
		},
		{
			Input:    "noooo",
			Expected: "n1o4",
		},
		{
			Input:    "Oooo",
			Expected: "O1o3",
		},
		{
			Input:    "anagram",
			Expected: "a1n1a1g1r1a1m1",
		},
		{
			Input:    "hello world",
			Expected: "h1e1l2o1w1o1r1l1d1",
		},
	}

	t.Run("Test Counter Letter", func(it *testing.T) {
		for _, tc := range testCases {

			res := CounterLetter(tc.Input)
			assert.Equal(it, tc.Expected, res)
		}
	})
}
