package main

func ConnectingValues(arr [][]int, val int, coord []int) [][]int {
	if len(arr) == 0 || len(arr[0]) == 0 {
		return [][]int{}
	}

	row, col := coord[0], coord[1]

	if row < 0 || row >= len(arr) || col < 0 || col >= len(arr[0]) {
		return [][]int{}
	}

	if arr[row][col] != val {
		return [][]int{}
	}

	var result [][]int

	type pos struct {
		r, c int
	}
	visited := make(map[pos]bool)

	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	queue := [][]int{{row, col}}
	visited[pos{row, col}] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		r, c := current[0], current[1]
		result = append(result, []int{r, c})

		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]

			if nr >= 0 && nr < len(arr) && nc >= 0 && nc < len(arr[0]) {
				p := pos{nr, nc}
				if !visited[p] && arr[nr][nc] == val {
					visited[p] = true
					queue = append(queue, []int{nr, nc})
				}
			}
		}
	}

	return result
}