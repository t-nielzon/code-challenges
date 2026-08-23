package main

func Solution(n int) int {
	if n <= 0 {
		return 0
	}

	sum := 0
	for i := 0; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}