package arraysandhashing_test

import (
	"testing"

	"github.com/lenardjombo/blind75/solutions/ArraysAndHashing"
)

func TestContainsDuplicates(t *testing.T){
	// Test cases definition
	testCases := []struct{
		name string
		numbers []int
		expected bool
	}{
		{
		name: "Array with no duplicates",
		numbers: []int{1,2,3,4,5},
		expected: false,
	},
	{
		name: "Array with duplicates",
		numbers: []int{1,2,3,2},
		expected:true,
	},
	{
		name: "Empty array",
		numbers: []int{},
		expected:false,
	},
	{
		name: "Array with one value",
		numbers: []int{99},
		expected:false,
	},
}

//Run each test case one by one
for _, tc := range testCases {
	t.Run(tc.name,func(t *testing.T) {
		result := arraysandhashing.ContainsDuplicate(tc.numbers)

		//Check if results == expected
		if result != tc.expected{
			t.Errorf("For numbers %v:\nGot %t\nBut expected %t",tc.numbers,result,tc.expected)
		}
	})
}
}
