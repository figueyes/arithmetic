package letters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "Y1e1s1",
			Expected: "Yes",
		},
		{
			Input:    "n1o4",
			Expected: "noooo",
		},
		{
			Input:    "O1o3",
			Expected: "Oooo",
		},
		{
			Input:    "y1e3s1",
			Expected: "yeees",
		},
	}

	t.Run("Decode Test", func(it *testing.T) {
		for _, tc := range testCases {
			res := Decode(tc.Input)
			assert.Equal(it, tc.Expected, res)
		}
	})
}
