package main

func MoveZeros(arr []int) []int {
	writeIdx := 0

	// move all non-zero elements to the front
	for readIdx := 0; readIdx < len(arr); readIdx++ {
		if arr[readIdx] != 0 {
			arr[writeIdx], arr[readIdx] = arr[readIdx], arr[writeIdx]
			writeIdx++
		}
	}

	// fill remaining positions with zeros
	for i := writeIdx; i < len(arr); i++ {
		arr[i] = 0
	}

	return arr
}