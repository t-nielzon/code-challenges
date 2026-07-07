func RideOfFortune(artifact []string, explorers []int) [][2]int {
	n := len(artifact)
	
	// Parse switches and their initial states
	switches := make(map[[2]int]byte)
	for i := 0; i < n; i++ {
		for j := 0; j < len(artifact[i]); j++ {
			if artifact[i][j] == 'A' || artifact[i][j] == 'B' {
				switches[[2]int{i, j}] = artifact[i][j]
			}
		}
	}
	
	results := make([][2]int, len(explorers))
	
	for idx, row := range explorers {
		pos := [2]int{row, -1}
		dir := [2]int{0, 1}
		
		for {
			pos[0] += dir[0]
			pos[1] += dir[1]
			
			if pos[1] < 0 {
				results[idx] = [2]int{-1, -1}
				break
			}
			if pos[1] >= n {
				results[idx] = [2]int{pos[0], n - 1}
				break
			}
			if pos[0] < 0 {
				results[idx] = [2]int{0, pos[1]}
				break
			}
			if pos[0] >= n {
				results[idx] = [2]int{n - 1, pos[1]}
				break
			}
			
			if state, exists := switches[pos]; exists {
				dir = routeCart(dir, state)
				switches[pos] = toggleState(state)
			}
		}
	}
	
	return results
}

func routeCart(dir [2]int, state byte) [2]int {
	if state == 'A' {
		if dir[0] == 0 && dir[1] == 1 {
			return [2]int{1, 0}
		} else if dir[0] == 0 && dir[1] == -1 {
			return [2]int{-1, 0}
		} else if dir[0] == 1 && dir[1] == 0 {
			return [2]int{0, 1}
		} else if dir[0] == -1 && dir[1] == 0 {
			return [2]int{0, -1}
		}
	} else {
		if dir[0] == 0 && dir[1] == 1 {
			return [2]int{-1, 0}
		} else if dir[0] == 0 && dir[1] == -1 {
			return [2]int{1, 0}
		} else if dir[0] == 1 && dir[1] == 0 {
			return [2]int{0, -1}
		} else if dir[0] == -1 && dir[1] == 0 {
			return [2]int{0, 1}
		}
	}
	return dir
}

func toggleState(state byte) byte {
	if state == 'A' {
		return 'B'
	}
	return 'A'
}