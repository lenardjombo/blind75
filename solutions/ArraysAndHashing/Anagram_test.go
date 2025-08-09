package arraysandhashing

import "testing"

func TestIsAnagram(t *testing.T) {
	testcases := []struct {
		name     string
		input1   string
		input2   string
		expected bool
	}{
		{
			name:     "empty strings",
			input1:   "",
			input2:   "",
			expected: true,
		},
		{
			name:     "simple anagram lowercase",
			input1:   "cat",
			input2:   "act",
			expected: true,
		},
		{
			name:     "different length",
			input1:   "not",
			input2:   "onto",
			expected: false,
		},
		{
			name:     "same letters different case",
			input1:   "Cat",
			input2:   "act",
			expected: false,
		},
		{
			name:     "anagram with spaces",
			input1:   "Dormitory",
			input2:   "Dirty Room",
			expected: false,
		},
		{
			name:     "exact same string",
			input1:   "hello",
			input2:   "hello",
			expected: true,
		},
		{
			name:     "completely different",
			input1:   "abc",
			input2:   "xyz",
			expected: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsAnagram(tc.input1, tc.input2)
			if actual != tc.expected {
				t.Errorf("IsAnagram(%q, %q) = %v; want %v",
					tc.input1, tc.input2, actual, tc.expected)
			}
		})
	}
}

