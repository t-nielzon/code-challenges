package kata

func PeakHeight(mountain []string) int {
	if len(mountain) == 0 {
		return 0
	}

	rows := len(mountain)
	// find max width
	maxCols := 0
	for _, row := range mountain {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}
	if maxCols == 0 {
		return 0
	}

	// build grid of heights: 0 for non-mountain, -1 for unresolved mountain
	height := make([][]int, rows)
	for r := 0; r < rows; r++ {
		height[r] = make([]int, maxCols)
		for c := 0; c < maxCols; c++ {
			if c < len(mountain[r]) && mountain[r][c] == '^' {
				height[r][c] = -1
			}
		}
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	maxHeight := 0

	for level := 1; ; level++ {
		// find all unresolved cells that border a non-unresolved cell (or edge)
		var toSet [][2]int
		for r := 0; r < rows; r++ {
			for c := 0; c < maxCols; c++ {
				if height[r][c] != -1 {
					continue
				}
				// check if on border or adjacent to resolved cell
				edge := r == 0 || r == rows-1 || c == 0 || c == maxCols-1
				if !edge {
					for _, d := range dirs {
						nr, nc := r+d[0], c+d[1]
						if nr >= 0 && nr < rows && nc >= 0 && nc < maxCols && height[nr][nc] != -1 {
							edge = true
							break
						}
					}
				}
				if edge {
					toSet = append(toSet, [2]int{r, c})
				}
			}
		}
		if len(toSet) == 0 {
			break
		}
		for _, pos := range toSet {
			height[pos[0]][pos[1]] = level
		}
		maxHeight = level
	}

	return maxHeight
}