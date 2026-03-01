package kata

func LargestSum(array []int) int {
	maxSum, currentSum := 0, 0
	for _, v := range array {
		currentSum += v
		if currentSum < 0 {
			currentSum = 0
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}