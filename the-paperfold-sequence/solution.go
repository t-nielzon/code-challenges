package kata

func PaperFold() func() int {
	n := 0
	return func() int {
		n++
		// Find the highest power of 2 that divides n
		// The position in the sequence is n (1-indexed)
		// We need to find k such that n = m * 2^k where m is odd
		// Then the value is determined by (m / 2) % 2
		// If ((n / 2^k) / 2) % 2 == 0, result is 1, else 0
		
		m := n
		for m&1 == 0 {
			m >>= 1
		}
		// m is now the odd part of n
		// The formula: if (m >> 1) & 1 == 0 then 1 else 0
		// Which simplifies to: 1 - ((m >> 1) & 1)
		return 1 - ((m >> 1) & 1)
	}
}