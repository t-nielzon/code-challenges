package main

func MaxSumSubarray(numbers []int) int {
	maxCurrent := 0
	maxGlobal := 0

	for _, num := range numbers {
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