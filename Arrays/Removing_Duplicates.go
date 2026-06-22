package Arrays

import "fmt"

func RemoveDup(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for i, value := range nums {
		if seen[value] == false {
			seen[value] = true

			result = append(result, nums[i])
		}
	}
	fmt.Print(len(result))
	return result

}
