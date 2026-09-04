package kata

func HexagonBeamMaxSum(n int, seq []int) int {
	// Build the hexagonal grid
	grid := make([][]int, 2*n-1)
	seqIdx := 0
	
	for i := 0; i < 2*n-1; i++ {
		var width int
		if i < n {
			width = n + i
		} else {
			width = 3*n - 2 - i
		}
		
		grid[i] = make([]int, width)
		for j := 0; j < width; j++ {
			grid[i][j] = seq[seqIdx%len(seq)]
			seqIdx++
		}
	}
	
	// Initialize with first horizontal beam
	maxSum := 0
	for j := 0; j < len(grid[0]); j++ {
		maxSum += grid[0][j]
	}
	
	// Check all horizontal beams (each row)
	for i := 1; i < 2*n-1; i++ {
		sum := 0
		for j := 0; j < len(grid[i]); j++ {
			sum += grid[i][j]
		}
		maxSum = max(maxSum, sum)
	}
	
	// Check all beams in diagonal direction 1 (row - col = constant)
	for k := -(2*n - 2); k <= 2*n-2; k++ {
		sum := 0
		for i := 0; i < 2*n-1; i++ {
			j := i - k
			if j >= 0 && j < len(grid[i]) {
				sum += grid[i][j]
			}
		}
		maxSum = max(maxSum, sum)
	}
	
	// Check all beams in diagonal direction 2 (row + col = constant)
	for k := 0; k <= 4*n-4; k++ {
		sum := 0
		for i := 0; i < 2*n-1; i++ {
			j := k - i
			if j >= 0 && j < len(grid[i]) {
				sum += grid[i][j]
			}
		}
		maxSum = max(maxSum, sum)
	}
	
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}