package kata

func PeakHeight(mountain []string) int {
	rows := len(mountain)
	if rows == 0 {
		return 0
	}

	cols := 0
	for _, r := range mountain {
		if len(r) > cols {
			cols = len(r)
		}
	}
	if cols == 0 {
		return 0
	}

	// grid[r][c] is true where there is a '^' remaining
	grid := make([][]bool, rows)
	remaining := 0
	for r := 0; r < rows; r++ {
		grid[r] = make([]bool, cols)
		for c := 0; c < cols; c++ {
			if c < len(mountain[r]) && mountain[r][c] == '^' {
				grid[r][c] = true
				remaining++
			}
		}
	}

	// a cell is on the current outer layer if it is a peak that neighbours a
	// non-peak cell or the edge of the grid (8-directional contact)
	isExposed := func(r, c int) bool {
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if dr == 0 && dc == 0 {
					continue
				}
				nr, nc := r+dr, c+dc
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					return true
				}
				if !grid[nr][nc] {
					return true
				}
			}
		}
		return false
	}

	height := 0
	for remaining > 0 {
		height++
		var peel [][2]int
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] && isExposed(r, c) {
					peel = append(peel, [2]int{r, c})
				}
			}
		}
		for _, p := range peel {
			grid[p[0]][p[1]] = false
			remaining--
		}
	}

	return height
}