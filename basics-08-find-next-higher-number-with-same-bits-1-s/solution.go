package main

func sameDigitsBitNumber(n int) int {
	c := n
	c0 := 0 // count of trailing zeros
	c1 := 0 // count of ones to the right of the rightmost non-trailing zero
	
	// count trailing zeros
	for (c & 1) == 0 && c != 0 {
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
	
	// set the rightmost non-trailing zero bit
	n |= (1 << pos)
	
	// clear all bits to the right of pos
	n &= ^((1 << pos) - 1)
	
	// insert (c1-1) ones on the right
	n |= (1 << (c1 - 1)) - 1
	
	return n
}