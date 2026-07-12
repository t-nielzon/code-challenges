package main

func RideOfFortune(artifact []string, explorers []int) [][]int {
	n := len(artifact)

	// Parse switches from artifact
	switches := make(map[[2]int]byte)
	for i, row := range artifact {
		for j, cell := range row {
			if cell == 'A' || cell == 'B' {
				switches[[2]int{i, j}] = byte(cell)
			}
		}
	}

	results := make([][]int, len(explorers))

	for idx, doorRow := range explorers {
		row, col := doorRow, -1
		dir := 0 // 0=east, 1=south, 2=west, 3=north

		for {
			// Move one step in current direction
			switch dir {
			case 0: // east
				col++
			case 1: // south
				row++
			case 2: // west
				col--
			case 3: // north
				row--
			}

			// Check if hit a wall
			if col == -1 || col == n || row == -1 || row == n {
				if col == -1 {
					results[idx] = []int{-1, -1}
				} else if col == n {
					results[idx] = []int{row, n - 1}
				} else if row == -1 {
					results[idx] = []int{0, col}
				} else { // row == n
					results[idx] = []int{n - 1, col}
				}
				break
			}

			// Check if hit a switch
			if state, ok := switches[[2]int{row, col}]; ok {
				dir = reroute(dir, state)
				switches[[2]int{row, col}] = toggle(state)
			}
		}
	}

	return results
}

func reroute(dir int, state byte) int {
	if state == 'A' {
		switch dir {
		case 0: // east -> south
			return 1
		case 1: // south -> east
			return 0
		case 2: // west -> north
			return 3
		case 3: // north -> west
			return 2
		}
	} else { // State B: opposite of A
		switch dir {
		case 0: // east -> north
			return 3
		case 1: // south -> west
			return 2
		case 2: // west -> south
			return 1
		case 3: // north -> east
			return 0
		}
	}
	return dir
}

func toggle(state byte) byte {
	if state == 'A' {
		return 'B'
	}
	return 'A'
}