package main

func createSpiral(N int) [][]int {
	if N < 1 {
		return [][]int{}
	}

	matrix := make([][]int, N)
	for i := range matrix {
		matrix[i] = make([]int, N)
	}

	num := 1
	top, bottom, left, right := 0, N-1, 0, N-1

	for top <= bottom && left <= right {
		// fill right
		for col := left; col <= right; col++ {
			matrix[top][col] = num
			num++
		}
		top++

		// fill down
		for row := top; row <= bottom; row++ {
			matrix[row][right] = num
			num++
		}
		right--

		// fill left
		if top <= bottom {
			for col := right; col >= left; col-- {
				matrix[bottom][col] = num
				num++
			}
			bottom--
		}

		// fill up
		if left <= right {
			for row := bottom; row >= top; row-- {
				matrix[row][left] = num
				num++
			}
			left++
		}
	}

	return matrix
}