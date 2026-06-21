// {1, 2, 1, 4, 5}
package Arrays

func SortedorNot(nums []int) bool {

	for i := 0; i < len(nums)-1; i++ {
		if !(nums[i] <= nums[i+1]) {

			return false

		}
	}
	return true
}
