package Arrays

func SecondLargest(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= len(nums)-1; j++ {
			if nums[i] > nums[j] && i > j {
				x := nums[i]
				nums[i] = nums[j]
				nums[j] = x
			}
		}

	}
	return nums[1]

}
