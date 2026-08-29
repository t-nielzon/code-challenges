package solution

import "math"

func HexagonBeamMaxSum(n int, seq []int) int {
	var grid [][]int
	seqIdx := 0
	
	// Build the hexagonal grid by filling with sequence values
	for i := 0; i < 2*n-1; i++ {
		var rowLen int
		if i < n {
			rowLen = n + i
		} else {
			rowLen = 3*n - 2 - i
		}
		
		row := make([]int, rowLen)
		for j := 0; j < rowLen; j++ {
			row[j] = seq[seqIdx%len(seq)]
			seqIdx++
		}
		grid = append(grid, row)
	}
	
	maxSum := math.MinInt
	
	// Get the visual offset (left indentation) for a given row
	getOffset := func(row int) int {
		if row < n {
			return n - 1 - row
		}
		return row - n + 1
	}
	
	// Get cell value at row with a given visual column
	getCell := func(row, visualCol int) (int, bool) {
		offset := getOffset(row)
		col := visualCol - offset
		if col >= 0 && col < len(grid[row]) {
			return grid[row][col], true
		}
		return 0, false
	}
	
	// Check all horizontal beams (each row)
	for i := 0; i < len(grid); i++ {
		sum := 0
		for j := 0; j < len(grid[i]); j++ {
			sum += grid[i][j]
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	
	// Check all diagonal beams in both directions
	for startRow := 0; startRow < len(grid); startRow++ {
		for startCol := 0; startCol < len(grid[startRow]); startCol++ {
			startVisualCol := getOffset(startRow) + startCol
			
			// Down-right diagonal (visual column increases by 1 per row)
			sum := 0
			for row := startRow; row < len(grid); row++ {
				visualCol := startVisualCol + (row - startRow)
				if val, ok := getCell(row, visualCol); ok {
					sum += val
				} else {
					break
				}
			}
			if sum > maxSum {
				maxSum = sum
			}
			
			// Down-left diagonal (visual column decreases by 1 per row)
			sum = 0
			for row := startRow; row < len(grid); row++ {
				visualCol := startVisualCol - (row - startRow)
				if val, ok := getCell(row, visualCol); ok {
					sum += val
				} else {
					break
				}
			}
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	
	return maxSum
}