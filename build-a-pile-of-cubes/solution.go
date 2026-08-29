package main

func findNb(m int64) int {
	var sum int64 = 0
	n := int64(0)
	for {
		n++
		sum += n * n * n
		if sum == m {
			return int(n)
		}
		if sum > m {
			return -1
		}
	}
}