func ModPow(x, y, n int64) int64 {
	if n == 1 {
		return 0
	}
	
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