package check

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckArithmeticSubarrays(t *testing.T) {
	testCases := []struct {
		nums, l, r []int
		expected   []bool
	}{
		{
			nums:     []int{4, 6, 5, 9, 3, 7},
			l:        []int{0, 0, 2},
			r:        []int{2, 3, 5},
			expected: []bool{true, false, true},
		},
		{
			nums:     []int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10},
			l:        []int{0, 1, 6, 4, 8, 7},
			r:        []int{4, 4, 9, 7, 9, 10},
			expected: []bool{false, true, false, false, true, true},
		},
		{
			nums:     []int{5, 8, 16, 24, 32},
			l:        []int{0, 1},
			r:        []int{3, 4},
			expected: []bool{false, true},
		},
		{
			nums:     []int{1},
			l:        []int{0},
			r:        []int{0},
			expected: []bool{false},
		},
	}

	for _, tc := range testCases {
		t.Run("CheckArithmeticSubarray", func(it *testing.T) {
			chas := CheckArithmeticSubarrays(tc.nums, tc.l, tc.r)
			assert.Equal(it, tc.expected, chas)
		})
	}
}
