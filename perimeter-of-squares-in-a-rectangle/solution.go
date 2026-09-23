package main

func perimeter(n int64) int64 {
	if n < 0 {
		return 0
	}
	
	var sum int64 = 1
	if n == 0 {
		return 4 * sum
	}
	
	a, b := int64(1), int64(1)
	sum += b
	
	for i := int64(2); i <= n; i++ {
		a, b = b, a+b
		sum += b
	}
	
	return 4 * sum
}