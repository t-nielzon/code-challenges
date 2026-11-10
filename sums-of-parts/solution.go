package kata

func PartsSums(ls []int) []int {
	total := 0
	for _, v := range ls {
		total += v
	}

	result := make([]int, len(ls)+1)
	result[0] = total

	sum := total
	for i, v := range ls {
		sum -= v
		result[i+1] = sum
	}

	return result
}