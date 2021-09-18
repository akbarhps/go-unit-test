package example

func Sum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

func Multiply(first int, second int) int {
	return first * second
}
