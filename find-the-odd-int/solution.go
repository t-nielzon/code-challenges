package main

func FindOdd(seq []int) int {
	result := 0
	for _, num := range seq {
		result ^= num
	}
	return result
}