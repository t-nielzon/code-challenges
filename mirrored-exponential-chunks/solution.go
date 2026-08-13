package main

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