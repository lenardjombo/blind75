package arraysandhashing

import (
	"github.com/lenardjombo/blind75/utils"
)

func ContainsDuplicate(nums []int) bool {
    seen := make(utils.IntSet)
    for _, value := range nums {
        if seen[value] {
            return true // duplicate found
        }
        seen[value] = true // mark as seen
    }
    return false // no duplicates
}
