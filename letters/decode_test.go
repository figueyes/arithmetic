package letters

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	testCases := []struct {
		input       string
		expected    string
		errExpected error
	}{
		{
			input:       "Y1e1s1",
			expected:    "Yes",
			errExpected: nil,
		},
		{
			input:       "n1o4",
			expected:    "noooo",
			errExpected: nil,
		},
		{
			input:       "O1o3",
			expected:    "Oooo",
			errExpected: nil,
		},
		{
			input:       "y1e3s1",
			expected:    "yeees",
			errExpected: nil,
		},
		{
			input:       "a1l1e1x",
			expected:    "",
			errExpected: errors.New("input has an odd length"),
		},
	}

	t.Run("Decode Test", func(it *testing.T) {
		for _, tc := range testCases {
			res, err := Decode(tc.input)
			assert.Equal(it, tc.errExpected, err)
			assert.Equal(it, tc.expected, res)
		}
	})
}
