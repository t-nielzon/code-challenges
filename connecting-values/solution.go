package kata

func ConnectingValues(arr [][]int, val int, coord [2]int) [][2]int {
	if len(arr) == 0 || coord[0] < 0 || coord[0] >= len(arr) || coord[1] < 0 || coord[1] >= len(arr[coord[0]]) {
		return nil
	}
	if arr[coord[0]][coord[1]] != val {
		return nil
	}

	rows := len(arr)
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, len(arr[i]))
	}

	var result [][2]int
	stack := [][2]int{coord}
	visited[coord[0]][coord[1]] = true

	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur)

		for _, d := range dirs {
			nr, nc := cur[0]+d[0], cur[1]+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < len(arr[nr]) && !visited[nr][nc] && arr[nr][nc] == val {
				visited[nr][nc] = true
				stack = append(stack, [2]int{nr, nc})
			}
		}
	}

	return result
}