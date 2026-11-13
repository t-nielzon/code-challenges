package main

func ConnectingValues(arr [][]int, val int, coord [2]int) [][2]int {
	rows := len(arr)
	if rows == 0 {
		return [][2]int{}
	}
	cols := len(arr[0])

	row, col := coord[0], coord[1]
	if row < 0 || row >= rows || col < 0 || col >= cols || arr[row][col] != val {
		return [][2]int{}
	}

	visited := make(map[[2]int]bool)
	result := [][2]int{}
	queue := [][2]int{{row, col}}
	visited[[2]int{row, col}] = true

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, dir := range directions {
			nr := current[0] + dir[0]
			nc := current[1] + dir[1]

			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !visited[[2]int{nr, nc}] && arr[nr][nc] == val {
				visited[[2]int{nr, nc}] = true
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}

	return result
}