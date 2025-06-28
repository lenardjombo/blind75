package arraysandhashing

type IntSet map[int]bool

func ContainsDuplicate(nums []int) bool {
    seen := make(IntSet)
    for _, value := range nums {
        if seen[value] {
            return true // duplicate found
        }
        seen[value] = true // mark as seen
    }
    return false // no duplicates
}
