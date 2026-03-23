package kata

func SplitAndAdd(arr []int, n int) []int {
	for i := 0; i < n && len(arr) > 1; i++ {
		mid := len(arr) / 2
		left := arr[:mid]
		right := arr[mid:]

		// right is always >= left in length; pad left by prepending zeros
		result := make([]int, len(right))
		copy(result, right)

		offset := len(right) - len(left)
		for j := 0; j < len(left); j++ {
			result[offset+j] += left[j]
		}

		arr = result
	}
	return arr
}