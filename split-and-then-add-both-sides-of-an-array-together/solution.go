package kata

func SplitAndAdd(arr []int, n int) []int {
	for i := 0; i < n && len(arr) > 1; i++ {
		arr = fold(arr)
	}
	return arr
}

func fold(arr []int) []int {
	length := len(arr)
	mid := length / 2
	
	// Split: left gets fewer elements when odd length
	left := arr[:mid]
	right := arr[mid:]
	
	// Right is always >= left in length
	result := make([]int, len(right))
	copy(result, right)
	
	// Add left to right, aligned to the right side
	offset := len(right) - len(left)
	for i := 0; i < len(left); i++ {
		result[offset+i] += left[i]
	}
	
	return result
}