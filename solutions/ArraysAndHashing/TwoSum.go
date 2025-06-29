package arraysandhashing

import(
	"github.com/lenardjombo/blind75/utils"
)

func TwoSum(nums []int, target int )([]int){
	seen := make(utils.IntMap)
	for index,value := range nums{
		complement := target - value
		if j,found := seen[complement];found{
			return []int{j,index}
		}
		seen[value] = index
	}
	return []int{}
}