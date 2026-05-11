package kata

func ConnectingValues(arr [][]int, val int, coord [2]int) [][2]int {
	result := [][2]int{}
	if len(arr) == 0 || len(arr[0]) == 0 {
		return result
	}
	rows := len(arr)
	cols := len(arr[0])
	r0, c0 := coord[0], coord[1]
	if r0 < 0 || r0 >= rows || c0 < 0 || c0 >= cols {
		return result
	}
	if arr[r0][c0] != val {
		return result
	}

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	stack := [][2]int{coord}
	visited[r0][c0] = true

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
			nr := cur[0] + d[0]
			nc := cur[1] + d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}
			if visited[nr][nc] {
				continue
			}
			if arr[nr][nc] != val {
				continue
			}
			visited[nr][nc] = true
			stack = append(stack, [2]int{nr, nc})
		}
	}

	return result
}