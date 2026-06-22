package Arrays

func LeftRotate(y []int) []int {
	n := len(y)
	y = append(y[1:n], y[0])
	return y
}
