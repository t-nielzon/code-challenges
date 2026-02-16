package kata

// sum of squares: 1² + 2² + ... + n² = n(n+1)(2n+1)/6
func Squares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}