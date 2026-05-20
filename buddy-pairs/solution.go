package kata

func sumProperDivisors(n uint64) uint64 {
	if n < 2 {
		return 0
	}
	sum := uint64(1)
	for i := uint64(2); i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			j := n / i
			if j != i {
				sum += j
			}
		}
	}
	return sum
}

func Buddy(start, limit uint64) []uint64 {
	for n := start; n <= limit; n++ {
		m := sumProperDivisors(n) - 1
		if m > n && sumProperDivisors(m) == n+1 {
			return []uint64{n, m}
		}
	}
	return nil
}