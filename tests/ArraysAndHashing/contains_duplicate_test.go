package arraysandhashing

import ("testing"
		"github.com/lenardjombo/blind75/solutions/ArraysAndHashing"
	)
func TestDuplicates(t *testing.T){
	input := []int{1,2,3,4,5,6,7,8,9}
	expected := false
	result := arraysandhashing.ContainsDuplicate(input)

	if result != expected{
		t.Errorf("Contains Duplicates (%v) = %v;want %v",input,result,expected)
	}
}