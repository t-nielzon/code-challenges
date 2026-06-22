package kata

func PartsSums(ls []int) []int {
	n := len(ls)
	result := make([]int, n+1)
	sum := 0
	// build from the end so each entry is the running suffix sum
	for i := n - 1; i >= 0; i-- {
		sum += ls[i]
		result[i] = sum
	}
	result[n] = 0
	return result
}