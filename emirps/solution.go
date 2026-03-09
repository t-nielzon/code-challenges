package kata

func FindEmirp(n int) [3]int {
	// sieve of eratosthenes up to a limit that covers reversed primes
	limit := n
	if limit < 10 {
		return [3]int{0, 0, 0}
	}

	// reversed primes can be larger than n, so we need a bigger sieve
	// a number below n has at most the same number of digits, so its
	// reverse is at most 10x larger
	sieveLimit := n * 10
	if sieveLimit < 1000 {
		sieveLimit = 1000
	}

	isPrime := make([]bool, sieveLimit+1)
	for i := 2; i <= sieveLimit; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= sieveLimit; i++ {
		if isPrime[i] {
			for j := i * i; j <= sieveLimit; j += i {
				isPrime[j] = false
			}
		}
	}

	count, largest, sum := 0, 0, 0

	for i := 13; i < n; i++ {
		if !isPrime[i] {
			continue
		}
		rev := reverse(i)
		if rev == i {
			continue
		}
		if rev <= sieveLimit && isPrime[rev] {
			count++
			largest = i
			sum += i
		}
	}

	return [3]int{count, largest, sum}
}

func reverse(n int) int {
	rev := 0
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev
}