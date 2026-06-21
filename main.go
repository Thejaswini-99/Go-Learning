package main

import (
	"fmt"
	"go-learning/Arrays"
)

func main() {
	nums := []int{3, 2, 1, 5, 6, 9, 8}
	SortedArray := []int{1, 2, 1, 4, 5}
	A := Arrays.Largest(nums)
	fmt.Println("Largest Element in the Array is", A)
	B := Arrays.SecondLargest(nums)
	fmt.Println("Second Largest Element in the Array is", B)
	C := Arrays.SortedorNot(SortedArray)
	fmt.Println("Sorted :", C)
}
