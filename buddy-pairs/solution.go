package kata

// sumProperDivisors returns the sum of proper divisors of n (divisors other than n).
func sumProperDivisors(n int) int {
	if n < 2 {
		return 0
	}
	// 1 always divides n and is a proper divisor for n > 1.
	sum := 1
	for i := 2; i*i <= n; i++ {
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

func Buddy(start, limit int) []int {
	for n := start; n <= limit; n++ {
		m := sumProperDivisors(n) - 1
		if m > n && sumProperDivisors(m) == n+1 {
			return []int{n, m}
		}
	}
	return nil
}