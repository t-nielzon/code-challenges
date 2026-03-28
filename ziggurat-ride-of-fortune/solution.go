package kata

func RideOfFortune(artifact []string, explorers []int) [][2]int {
	n := len(artifact)
	grid := make([][]byte, n)
	for i := range artifact {
		grid[i] = []byte(artifact[i])
	}

	results := make([][2]int, len(explorers))

	for ei, door := range explorers {
		r, c := door, 0
		dr, dc := 0, 1

		for {
			ch := grid[r][c]
			if ch == 'A' || ch == 'B' {
				if ch == 'A' {
					dr, dc = dc, dr
				} else {
					dr, dc = -dc, -dr
				}
				grid[r][c] ^= 'A' ^ 'B'
			}

			nr, nc := r+dr, c+dc
			if nr < 0 || nr >= n || nc < 0 || nc >= n {
				if nc < 0 {
					results[ei] = [2]int{-1, -1}
				} else {
					results[ei] = [2]int{r, c}
				}
				break
			}
			r, c = nr, nc
		}
	}

	return results
}