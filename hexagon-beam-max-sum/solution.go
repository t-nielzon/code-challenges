func MaxHexagonBeam(n int, seq []int) int {
	grid := buildGrid(n, seq)
	maxSum := -(1 << 31)

	// horizontal beams (rows)
	for _, row := range grid {
		sum := 0
		for _, val := range row {
			sum += val
		}
		if sum > maxSum {
			maxSum = sum
		}
	}

	// vertical beams (columns)
	totalRows := 2*n - 1
	maxCols := 0
	for _, row := range grid {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	for col := 0; col < maxCols; col++ {
		sum := 0
		hasCell := false
		for row := 0; row < totalRows; row++ {
			if col < len(grid[row]) {
				sum += grid[row][col]
				hasCell = true
			}
		}
		if hasCell && sum > maxSum {
			maxSum = sum
		}
	}

	// diagonal beams going from upper-left to bottom-right
	// pattern: col = start_col + min(row, totalRows - 1 - row)
	for start_col := 0; start_col < n; start_col++ {
		sum := 0
		hasCell := false
		for row := 0; row < totalRows; row++ {
			col := start_col + minInt(row, totalRows-1-row)
			if col < len(grid[row]) {
				sum += grid[row][col]
				hasCell = true
			}
		}
		if hasCell && sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func buildGrid(n int, seq []int) [][]int {
	var grid [][]int
	seqIndex := 0
	totalRows := 2*n - 1

	for row := 0; row < totalRows; row++ {
		var r []int
		var rowLen int

		if row < n {
			rowLen = n + row
		} else {
			rowLen = 3*n - 2 - row
		}

		for col := 0; col < rowLen; col++ {
			r = append(r, seq[seqIndex%len(seq)])
			seqIndex++
		}

		grid = append(grid, r)
	}

	return grid
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}