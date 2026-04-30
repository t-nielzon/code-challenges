package kata

func sumProperDivisors(n int) int {
	if n < 2 {
		return 0
	}
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum
}

func Buddy(start, limit int) []int {
	for n := start; n <= limit; n++ {
		m := sumProperDivisors(n) - 1
		if m > n && sumProperDivisors(m) == n+1 {
			return []int{n, m}
		}
	}
	return nil
}