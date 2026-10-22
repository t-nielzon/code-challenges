package kata

func MirroredExponentialChunks(arr []int) [][]int {
	if len(arr) == 0 {
		return [][]int{}
	}
	
	n := len(arr)
	var chunkSizes []int
	
	remaining := n
	
	// center chunk of size 1 (2^0)
	chunkSizes = append(chunkSizes, 1)
	remaining--
	
	// add layers symmetrically around center
	layer := 1
	for remaining > 0 {
		size := 1 << uint(layer) // 2^layer
		
		// add full symmetric layer only if we have enough elements for both sides
		if 2*size <= remaining {
			chunkSizes = append([]int{size}, chunkSizes...)
			chunkSizes = append(chunkSizes, size)
			remaining -= 2 * size
			layer++
		} else {
			break
		}
	}
	
	// add any remaining elements as final chunk
	if remaining > 0 {
		chunkSizes = append(chunkSizes, remaining)
	}
	
	// create chunks from sizes
	var result [][]int
	idx := 0
	for _, size := range chunkSizes {
		result = append(result, arr[idx:idx+size])
		idx += size
	}
	
	return result
}