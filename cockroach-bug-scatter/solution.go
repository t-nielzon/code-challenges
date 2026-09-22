package main

func CockroachCount(grid []string) []int {
	result := make([]int, 10)
	
	// Parse grid to find all cockroaches
	var cockroaches []struct {
		row, col int
		dir      rune
	}
	
	for r, line := range grid {
		for c, ch := range line {
			if ch == 'U' || ch == 'D' || ch == 'L' || ch == 'R' {
				cockroaches = append(cockroaches, struct {
					row, col int
					dir      rune
				}{r, c, ch})
			}
		}
	}
	
	// Simulate each cockroach
	for _, roach := range cockroaches {
		r, c := roach.row, roach.col
		dir := roach.dir
		
		// Move until finding a hole
		maxIter := 10000
		for i := 0; i < maxIter; i++ {
			// Calculate next position
			var nr, nc int
			switch dir {
			case 'U':
				nr, nc = r-1, c
			case 'D':
				nr, nc = r+1, c
			case 'L':
				nr, nc = r, c-1
			case 'R':
				nr, nc = r, c+1
			}
			
			// Check bounds
			if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[nr]) {
				dir = turnLeft(dir)
				continue
			}
			
			ch := grid[nr][nc]
			
			// Check if next cell is a wall
			if ch == '+' || ch == '|' || ch == '-' {
				dir = turnLeft(dir)
			} else if ch >= '0' && ch <= '9' {
				// Found a hole
				holeNum := int(ch - '0')
				result[holeNum]++
				break
			} else {
				// Move forward
				r, c = nr, nc
			}
		}
	}
	
	return result
}

func turnLeft(dir rune) rune {
	switch dir {
	case 'U':
		return 'L'
	case 'L':
		return 'D'
	case 'D':
		return 'R'
	case 'R':
		return 'U'
	}
	return dir
}