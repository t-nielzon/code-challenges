package kata

func LargestSum(arr []int) int {
	maxSum := 0
	currentSum := 0
	for _, v := range arr {
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