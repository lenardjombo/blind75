package arraysandhashing

import (
	"testing"

	"github.com/lenardjombo/blind75/solutions/ArraysAndHashing"
)

func TestTwoSum(t *testing.T) {
	testcases := []struct {
		name     string
		input    []int
		target   int
		expected []int
	}{
		{
			name:     "Two sum basic case",
			input:    []int{3, 2, 5},
			target:   7,
			expected: []int{1, 2}, // 2 + 5 = 7
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			result := arraysandhashing.TwoSum(tc.input, tc.target)

			if len(result) != len(tc.expected) {
				t.Errorf("Expected length %d, got %d", len(tc.expected), len(result))
				return
			}

			for i := range result {
				if result[i] != tc.expected[i] {
					t.Errorf("At index %d: expected %d, got %d", i, tc.expected[i], result[i])
				}
			}
		})
	}
}
