package kata

func Finance(n int) int64 {
	// each cell (w,d) where 0 ≤ w ≤ d ≤ n contributes w+d, which sums to n(n+1)(n+2)/2
	m := int64(n)
	return m * (m + 1) * (m + 2) / 2
}