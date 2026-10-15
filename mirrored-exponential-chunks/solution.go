package kata

func ChunkExponentially(arr []int) [][]int {
	n := len(arr)
	if n == 0 {
		return [][]int{}
	}

	var sizes []int
	total := 0

	for d := 0; total < n; d++ {
		var need int
		if d == 0 {
			need = 1
		} else {
			need = 2 * (1 << uint(d))
		}

		if total+need <= n {
			if d == 0 {
				sizes = append(sizes, 1)
			} else {
				sizes = append(sizes, 1<<uint(d))
			}
			total += need
		} else {
			break
		}
	}

	var result [][]int
	idx := 0
	numDist := len(sizes)

	for d := numDist - 1; d >= 1; d-- {
		size := sizes[d]
		chunk := arr[idx : idx+size]
		result = append(result, chunk)
		idx += size
	}

	if numDist > 0 {
		chunk := arr[idx : idx+1]
		result = append(result, chunk)
		idx += 1
	}

	for d := 1; d < numDist; d++ {
		size := sizes[d]
		chunk := arr[idx : idx+size]
		result = append(result, chunk)
		idx += size
	}

	return result
}