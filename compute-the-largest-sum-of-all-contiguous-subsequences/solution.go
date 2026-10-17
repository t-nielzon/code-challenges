package main

func largestSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxCurrent := 0
	maxGlobal := 0

	for _, num := range arr {
		maxCurrent = maxCurrent + num
		if maxCurrent < 0 {
			maxCurrent = 0
		}
		if maxGlobal < maxCurrent {
			maxGlobal = maxCurrent
		}
	}

	return maxGlobal
}