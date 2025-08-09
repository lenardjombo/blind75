package arraysandhashing

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	testcases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "has duplicates",
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: false,
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: false,
		},
		{
			name:     "all elements same",
			input:    []int{7, 7, 7, 7},
			expected: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := ContainsDuplicate(tc.input)
			if got != tc.expected {
				t.Errorf("ContainsDuplicate(%v) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}

