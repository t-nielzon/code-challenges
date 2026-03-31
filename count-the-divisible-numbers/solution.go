package kata

func Divisors(x, y, k int) int {
	return y/k - (x-1)/k
}