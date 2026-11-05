package main

func PascalsDiagonal(n, l int) []int {
	if l == 0 {
		return []int{}
	}

	result := make([]int, l)
	result[0] = 1

	for i := 1; i < l; i++ {
		result[i] = result[i-1] * (n + i) / i
	}

	return result
}