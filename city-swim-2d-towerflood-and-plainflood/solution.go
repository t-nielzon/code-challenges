package main

func GetWater(towers []int) int {
	if len(towers) <= 2 {
		return 0
	}

	n := len(towers)
	maxLeft := make([]int, n)
	maxRight := make([]int, n)

	// calculate max height from start to each position
	maxLeft[0] = towers[0]
	for i := 1; i < n; i++ {
		if towers[i] > maxLeft[i-1] {
			maxLeft[i] = towers[i]
		} else {
			maxLeft[i] = maxLeft[i-1]
		}
	}

	// calculate max height from each position to end
	maxRight[n-1] = towers[n-1]
	for i := n - 2; i >= 0; i-- {
		if towers[i] > maxRight[i+1] {
			maxRight[i] = towers[i]
		} else {
			maxRight[i] = maxRight[i+1]
		}
	}

	// calculate trapped water at each position
	trapped := 0
	for i := 0; i < n; i++ {
		waterLevel := maxLeft[i]
		if maxRight[i] < waterLevel {
			waterLevel = maxRight[i]
		}
		trapped += waterLevel - towers[i]
	}

	return trapped
}