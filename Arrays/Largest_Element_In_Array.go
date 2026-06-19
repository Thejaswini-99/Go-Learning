// Find the largest element in the array
package Arrays

func Largest(nums []int) int {
	max := nums[0]
	for _, value := range nums {
		if max < value {
			max = value
		}
	}
	return max
}
