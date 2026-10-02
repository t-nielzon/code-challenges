package main

func FindWater(h []int) int {
	if len(h) == 0 {
		return 0
	}

	n := len(h)
	maxLeft := make([]int, n)
	maxRight := make([]int, n)

	// precompute maximum height to the left of each position
	maxLeft[0] = h[0]
	for i := 1; i < n; i++ {
		maxLeft[i] = max(maxLeft[i-1], h[i])
	}

	// precompute maximum height to the right of each position
	maxRight[n-1] = h[n-1]
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], h[i])
	}

	// calculate trapped water: water level at position i is the minimum
	// of the max heights on both sides, minus the tower height at i
	water := 0
	for i := 0; i < n; i++ {
		waterLevel := min(maxLeft[i], maxRight[i])
		water += waterLevel - h[i]
	}

	return water
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}