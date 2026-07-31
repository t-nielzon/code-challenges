package kata

func largestSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxCurrent := 0
	maxGlobal := 0

	for _, num := range nums {
		maxCurrent = max(0, maxCurrent+num)
		maxGlobal = max(maxGlobal, maxCurrent)
	}

	return maxGlobal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}