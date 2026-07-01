package kata

func FindEmirp(n int) [3]int {
	if n < 2 {
		return [3]int{0, 0, 0}
	}

	// sieve of eratosthenes for primes below n (exclusive upper limit)
	limit := n - 1
	isComposite := make([]bool, limit+1)
	for i := 2; i*i <= limit; i++ {
		if !isComposite[i] {
			for j := i * i; j <= limit; j += i {
				isComposite[j] = true
			}
		}
	}

	primeSet := make(map[int]bool)
	for i := 2; i <= limit; i++ {
		if !isComposite[i] {
			primeSet[i] = true
		}
	}

	count, largest, sum := 0, 0, 0
	for p := range primeSet {
		r := reverse(p)
		// emirp: reversed is a different prime, so palindromes are excluded
		if r != p && primeSet[r] {
			count++
			sum += p
			if p > largest {
				largest = p
			}
		}
	}

	return [3]int{count, largest, sum}
}

func reverse(num int) int {
	rev := 0
	for num > 0 {
		rev = rev*10 + num%10
		num /= 10
	}
	return rev
}