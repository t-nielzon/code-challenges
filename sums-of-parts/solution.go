package kata

func PartsSums(ls []int) []int {
	result := make([]int, len(ls)+1)
	sum := 0
	for i := len(ls) - 1; i >= 0; i-- {
		sum += ls[i]
		result[i] = sum
	}
	result[len(ls)] = 0
	return result
}