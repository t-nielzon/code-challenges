package main

func MoveZeros(arr []int) []int {
	result := make([]int, len(arr))
	index := 0

	for _, val := range arr {
		if val != 0 {
			result[index] = val
			index++
		}
	}

	for i := index; i < len(arr); i++ {
		result[i] = 0
	}

	return result
}