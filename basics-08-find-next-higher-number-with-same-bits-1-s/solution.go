package kata

func NextHigherNumberWithSameBits(num int) int {
	c := num
	c0 := 0 // count of trailing zeros
	c1 := 0 // count of ones after trailing zeros
	
	// count trailing zeros
	for ((c & 1) == 0) && c != 0 {
		c0++
		c >>= 1
	}
	
	// count ones after trailing zeros
	for (c & 1) == 1 {
		c1++
		c >>= 1
	}
	
	// position of rightmost non-trailing zero
	pos := c0 + c1
	
	// flip the rightmost non-trailing zero to 1
	num |= (1 << uint(pos))
	
	// clear all bits to the right of pos
	num &= ^((1 << uint(pos)) - 1)
	
	// insert (c1-1) ones on the right to keep same bit count
	num |= (1 << uint(c1-1)) - 1
	
	return num
}