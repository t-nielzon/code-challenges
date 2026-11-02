package main

func Modpow(x, y, n int64) int64 {
	result := int64(1)
	x = x % n
	
	for y > 0 {
		if y%2 == 1 {
			result = (result * x) % n
		}
		y = y >> 1
		x = (x * x) % n
	}
	
	return result
}