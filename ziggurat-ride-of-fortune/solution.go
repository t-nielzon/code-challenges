package solution

func RideOfFortune(artifact []string, explorers []int) [][2]int {
	n := len(artifact)
	grid := make([][]byte, n)
	for i, row := range artifact {
		grid[i] = []byte(row)
	}

	dr := [4]int{-1, 0, 1, 0}
	dc := [4]int{0, 1, 0, -1}

	routeA := [4]int{3, 2, 1, 0}
	routeB := [4]int{1, 0, 3, 2}

	result := make([][2]int, len(explorers))

	for i, door := range explorers {
		r, c := door, -1
		dir := 1

		for {
			nr, nc := r+dr[dir], c+dc[dir]
			if nc < 0 {
				result[i] = [2]int{-1, -1}
				break
			}
			if nr < 0 || nr >= n || nc >= n {
				result[i] = [2]int{r, c}
				break
			}
			r, c = nr, nc
			cell := grid[r][c]
			if cell == 'A' {
				dir = routeA[dir]
				grid[r][c] = 'B'
			} else if cell == 'B' {
				dir = routeB[dir]
				grid[r][c] = 'A'
			}
		}
	}

	return result
}