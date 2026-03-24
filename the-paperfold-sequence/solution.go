package kata

func PaperFold() func() int {
	// use the bit-trick: the value at 1-indexed position n is determined by
	// removing trailing zeros from n, then checking the bit above them
	n := 0
	return func() int {
		n++
		// find the highest power of 2 dividing n
		v := n
		for v&1 == 0 {
			v >>= 1
		}
		// the result depends on v mod 4
		if v%4 == 1 {
			return 1
		}
		return 0
	}
}