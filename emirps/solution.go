package kata

func FindEmirp(n int) [3]int {
	if n <= 12 {
		return [3]int{0, 0, 0}
	}

	// sieve of eratosthenes up to a limit that covers reversed primes
	// reversed primes can be at most same digit count, so we need sieve
	// up to the largest possible reverse of numbers below n
	limit := n
	// ensure sieve is large enough for reversed numbers
	// a number below n has at most the same number of digits, so its
	// reverse is bounded by 10^(digits) - 1
	power := 1
	for power < n {
		power *= 10
	}
	if power-1 > limit {
		limit = power - 1
	}

	sieve := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		sieve[i] = true
	}
	for i := 2; i*i <= limit; i++ {
		if sieve[i] {
			for j := i * i; j <= limit; j += i {
				sieve[j] = false
			}
		}
	}

	count := 0
	largest := 0
	sum := 0

	for i := 13; i < n; i++ {
		if !sieve[i] {
			continue
		}
		rev := reverseNum(i)
		if rev == i {
			continue
		}
		if rev <= limit && sieve[rev] {
			count++
			largest = i
			sum += i
		}
	}

	return [3]int{count, largest, sum}
}

func reverseNum(n int) int {
	rev := 0
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev
}