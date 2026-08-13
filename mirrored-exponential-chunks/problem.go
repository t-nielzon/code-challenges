package main

/*
Mirrored Exponential Chunks

Using a single function provided with an array, create a new one in the following ways:
- Separate array elements into chunks. Each chunk will contain 2^abs(n) elements where n is the "distance" to the median chunk.
- The median chunk(s) will have one single element, or will not appear in the output if there is some ambiguity about it.
- If an outlying chunk does not have enough elements to reach the 2^abs(n) requirement, it will contain as many elements as remain.
- Chunks in mirrored places (compared to the median chunk) should be of the same length.
- You must maintain original element order within each chunk.
- You must maintain original element order from chunk to chunk.
*/

func MirroredExponentialChunks(arr []int) [][]int {
	n := len(arr)
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{arr[0]}}
	}
	
	maxPower := 0
	for (1 << uint(maxPower)) < n {
		maxPower++
	}
	
	layers := make([]int, 0)
	remaining := n
	for power := maxPower; power >= 0; power-- {
		chunkSize := 1 << uint(power)
		if remaining >= 2*chunkSize+1 {
			layers = append(layers, chunkSize)
			remaining -= 2 * chunkSize
		}
	}
	
	if remaining != 1 || len(layers) == 0 {
		return [][]int{arr}
	}
	
	result := make([][]int, 0, 2*len(layers)+1)
	left := 0
	
	for _, chunkSize := range layers {
		result = append(result, arr[left:left+chunkSize])
		left += chunkSize
	}
	
	result = append(result, arr[left:left+1])
	left++
	
	for i := len(layers) - 1; i >= 0; i-- {
		chunkSize := layers[i]
		result = append(result, arr[left:left+chunkSize])
		left += chunkSize
	}
	
	return result
}