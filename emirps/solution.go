package kata

func FindEmirp(n int) [3]int {
	// sieve large enough to cover reversed primes
	limit := n
	if limit < 10 {
		return [3]int{0, 0, 0}
	}

	// upper bound for reversed numbers: a number below n can reverse to
	// something with the same digit count, so we need a sieve that covers
	// the largest possible reversal
	sieveSize := 10
	for sieveSize < n*10 {
		sieveSize *= 10
	}

	sieve := make([]bool, sieveSize)
	for i := 2; i < sieveSize; i++ {
		sieve[i] = true
	}
	for i := 2; i*i < sieveSize; i++ {
		if sieve[i] {
			for j := i * i; j < sieveSize; j += i {
				sieve[j] = false
			}
		}
	}

	count, largest, sum := 0, 0, 0

	for i := 13; i < n; i++ {
		if !sieve[i] {
			continue
		}
		rev := reverseNum(i)
		if rev == i {
			continue
		}
		if rev < sieveSize && sieve[rev] {
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