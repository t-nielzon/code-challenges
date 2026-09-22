func SameNumberOfBits(n int) int {
	c := n
	c0 := 0
	c1 := 0
	
	// count trailing zeros (rightmost zeros with no ones to their right)
	for ((c & 1) == 0) && c != 0 {
		c0++
		c >>= 1
	}
	
	// count consecutive ones after the trailing zeros
	for (c & 1) == 1 {
		c1++
		c >>= 1
	}
	
	// to get next higher number: flip the rightmost 0 that has 1s to its right
	pos := c0 + c1
	n |= (1 << pos)
	
	// clear everything to the right of the flipped bit to prepare for rearranging
	n &= ^((1 << pos) - 1)
	
	// place ones as far right as possible to minimize the result (all zeros left, all ones right)
	n |= (1 << (c1 - 1)) - 1
	
	return n
}