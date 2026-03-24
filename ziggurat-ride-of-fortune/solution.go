package kata

func RideOfFortune(artifact []string, explorers []int) [][2]int {
	n := len(artifact)

	// parse switches into a mutable grid
	grid := make([][]byte, n)
	for i, row := range artifact {
		grid[i] = []byte(row)
	}

	// directions: 0=north, 1=east, 2=south, 3=west
	// state A routing: entering moving west->north, east->south, south->east, north->west
	// "entering moving west" means cart direction is east (moving east into switch from west side)
	// Re-read: "a cart enters by moving west" means the cart is moving westward.
	// State A: moving west->routed north, moving east->routed south, moving south->routed east, moving north->routed west
	// State B: opposite orthogonal: moving west->south, moving east->north, moving south->west, moving north->east

	// direction deltas: 0=north(-1,0), 1=east(0,+1), 2=south(+1,0), 3=west(0,-1)
	dr := [4]int{-1, 0, 1, 0}
	dc := [4]int{0, 1, 0, -1}

	// State A: map incoming direction to outgoing direction
	// west(3)->north(0), east(1)->south(2), south(2)->east(1), north(0)->west(3)
	routeA := [4]int{3, 2, 1, 0}
	// State B: opposite orthogonal
	// west(3)->south(2), east(1)->north(0), south(2)->west(3), north(0)->east(1)
	routeB := [4]int{1, 0, 3, 2}

	results := make([][2]int, len(explorers))

	for idx, door := range explorers {
		// explorer enters from west wall at row=door, moving east
		r, c := door, 0
		dir := 1 // east

		// if starting cell is a switch, handle it; otherwise move
		for {
			// check if current cell is a switch
			if grid[r][c] == 'A' || grid[r][c] == 'B' {
				var newDir int
				if grid[r][c] == 'A' {
					newDir = routeA[dir]
				} else {
					newDir = routeB[dir]
				}
				// toggle switch
				if grid[r][c] == 'A' {
					grid[r][c] = 'B'
				} else {
					grid[r][c] = 'A'
				}
				dir = newDir
			}

			// move in current direction
			nr, nc := r+dr[dir], c+dc[dir]

			// check if out of bounds (hit a wall)
			if nr < 0 || nr >= n || nc < 0 || nc >= n {
				// hit a wall
				if dir == 3 {
					// west wall -> back outside
					results[idx] = [2]int{-1, -1}
				} else {
					// north, east, or south wall -> portal
					results[idx] = [2]int{nr, nc}
				}
				break
			}

			r, c = nr, nc
		}
	}

	return results
}