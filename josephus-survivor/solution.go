package main

func Josephus(n int, k int) int {
	result := 0
	for i := 2; i <= n; i++ {
		result = (result + k) % i
	}
	return result + 1
}