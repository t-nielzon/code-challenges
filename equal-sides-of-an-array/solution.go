package kata

func FindEvenIndex(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}

	left := 0
	for i, v := range arr {
		if left == total-left-v {
			return i
		}
		left += v
	}

	return -1
}