package kata

func MirroredExponentialChunks(array []interface{}) [][]interface{} {
	if len(array) == 0 {
		return [][]interface{}{}
	}

	n := len(array)

	// Calculate chunk sizes that form symmetric pairs
	var chunkSizes []int
	total := 1 // center chunk count
	size := 1
	for total+size*2 <= n {
		size *= 2
		chunkSizes = append(chunkSizes, size)
		total += size * 2
	}

	result := make([][]interface{}, 0)
	pos := 0

	// Left chunks (from largest to smallest)
	for i := len(chunkSizes) - 1; i >= 0; i-- {
		sz := chunkSizes[i]
		chunk := make([]interface{}, 0)
		for j := 0; j < sz && pos < n; j++ {
			chunk = append(chunk, array[pos])
			pos++
		}
		if len(chunk) > 0 {
			result = append(result, chunk)
		}
	}

	// Center chunk
	if pos < n {
		chunk := []interface{}{array[pos]}
		result = append(result, chunk)
		pos++
	}

	// Right chunks (from smallest to largest)
	for i := 0; i < len(chunkSizes); i++ {
		sz := chunkSizes[i]
		chunk := make([]interface{}, 0)
		for j := 0; j < sz && pos < n; j++ {
			chunk = append(chunk, array[pos])
			pos++
		}
		if len(chunk) > 0 {
			result = append(result, chunk)
		}
	}

	return result
}