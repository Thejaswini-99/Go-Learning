package Arrays

func SecondLargest(nums []int) int {
	Largest := nums[0]
	SecondLargest := -1

	for _, value := range nums {
		if Largest < value {
			SecondLargest = Largest
			Largest = value
		} else if value < Largest && value > SecondLargest && value != Largest {
			SecondLargest = value
		}

	}
	if SecondLargest == -1 {
		return -1
	}

	return SecondLargest

}
