package main

import (
	"fmt"
	"go-learning/Arrays"
)

func main() {
	Arr := []int{1, 2, 3, 4, 5}
	nums := []int{3, 2, 1, 5, 6, 9, 8}
	SortedArray := []int{0, 0, 3, 3, 5, 6}
	A := Arrays.Largest(nums)
	fmt.Println("Largest Element in the Array is", A)
	B := Arrays.SecondLargest(nums)
	fmt.Println("Second Largest Element in the Array is", B)
	C := Arrays.SortedorNot(SortedArray)
	fmt.Println("Sorted :", C)
	D := Arrays.RemoveDup(SortedArray)
	fmt.Println("Unique Array:", D)
	E := Arrays.LeftRotate(Arr)
	fmt.Println("Left Rotate Array by 1 is", E)
}
