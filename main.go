package main

import (
	"fmt"
	"go-learning/Arrays"
)

func main() {
	nums := []int{3, 2, 1, 5, 6, 9, 8}
	x := Arrays.Largest(nums)
	fmt.Println("Largest Element in the Array is", x)
	y := Arrays.SecondLargest(nums)
	fmt.Println("Second Largest Element in the Array is", y)
}
