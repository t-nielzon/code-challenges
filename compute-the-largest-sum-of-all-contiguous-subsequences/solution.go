package kata

func LargestSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxSum := 0
	currentSum := 0

	for _, num := range arr {
		currentSum += num
		if currentSum < 0 {
			currentSum = 0
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}