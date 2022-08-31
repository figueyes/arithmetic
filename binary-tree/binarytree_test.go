package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintTree(t *testing.T) {
	testCases := []struct {
		array    []int
		expected []string
	}{
		{
			array: []int{
				10, 7, 19, 5, 8, 34, 21, 2, 14, 3, 16, 11, 30, 26, 1,
			},
			expected: []string{"1", "2", "3", "5", "7", "8", "10", "11", "14", "16", "19", "21", "26", "30", "34"},
		},
		{
			array:    nil,
			expected: []string{"0"},
		},
		{
			array:    []int{1, 72, 34},
			expected: []string{"1", "34", "72"},
		},
	}

	for _, tc := range testCases {
		t.Run("print ordered tree Node", func(it *testing.T) {
			node := &Node{}
			for _, value := range tc.array {
				node.Add(value)
			}
			res := node.String(node)
			assert.Equal(it, tc.expected, res)
		})
	}
}
