package kata

func Cycle(n int) int {
	if n%2 == 0 || n%5 == 0 {
		return -1
	}
	r := 1 % n
	for i := 1; i <= n; i++ {
		r = (r * 10) % n
		if r == 1 {
			return i
		}
	}
	return -1
}