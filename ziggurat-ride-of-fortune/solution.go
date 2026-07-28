package main

func RideOfFortune(artifact []string, explorers []int) [][2]int {
	n := len(artifact)

	switches := make(map[[2]int]rune)
	for i := 0; i < n; i++ {
		for j := 0; j < len(artifact[i]); j++ {
			if artifact[i][j] == 'A' || artifact[i][j] == 'B' {
				switches[[2]int{i, j}] = artifact[i][j]
			}
		}
	}

	var results [][2]int
	for _, doorRow := range explorers {
		exit := simulatePath(doorRow, n, switches)
		results = append(results, exit)
	}

	return results
}

func simulatePath(doorRow int, n int, switches map[[2]int]rune) [2]int {
	row, col := doorRow, 0
	dir := 0 // 0=East, 1=South, 2=West, 3=North

	deltas := [][2]int{
		{0, 1},   // East
		{1, 0},   // South
		{0, -1},  // West
		{-1, 0},  // North
	}

	for {
		dr, dc := deltas[dir][0], deltas[dir][1]
		newRow := row + dr
		newCol := col + dc

		if newRow < 0 || newRow >= n || newCol < 0 || newCol >= n {
			if newCol < 0 {
				return [2]int{-1, -1}
			} else if newCol >= n {
				return [2]int{row, n - 1}
			} else if newRow < 0 {
				return [2]int{0, col}
			} else if newRow >= n {
				return [2]int{n - 1, col}
			}
		}

		row, col = newRow, newCol

		if switchState, exists := switches[[2]int{row, col}]; exists {
			if switchState == 'A' {
				dir = directionForStateA(dir)
			} else {
				dir = directionForStateB(dir)
			}
			switches[[2]int{row, col}] = toggleSwitch(switchState)
		}
	}
}

func directionForStateA(dir int) int {
	switch dir {
	case 0:
		return 1 // East -> South
	case 1:
		return 0 // South -> East
	case 2:
		return 3 // West -> North
	case 3:
		return 2 // North -> West
	}
	return dir
}

func directionForStateB(dir int) int {
	switch dir {
	case 0:
		return 3 // East -> North
	case 1:
		return 2 // South -> West
	case 2:
		return 1 // West -> South
	case 3:
		return 0 // North -> East
	}
	return dir
}

func toggleSwitch(state rune) rune {
	if state == 'A' {
		return 'B'
	}
	return 'A'
}