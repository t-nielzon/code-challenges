package main

func Josephus(n, k int) int {
	pos := 0
	for i := 2; i <= n; i++ {
		pos = (pos + k) % i
	}
	return pos + 1
}