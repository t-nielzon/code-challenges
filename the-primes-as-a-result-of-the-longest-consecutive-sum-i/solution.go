package kata

func LongestConsecPrimeSum(n int) []int {
	// sieve of eratosthenes up to n
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	// collect primes below n
	var primes []int
	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	// prefix sums for sliding window
	prefix := make([]int, len(primes)+1)
	for i, p := range primes {
		prefix[i+1] = prefix[i] + p
	}

	bestLen := 0
	var result []int

	// try all window lengths from largest down
	for length := len(primes); length >= 1; length-- {
		if length < bestLen {
			break
		}
		for start := 0; start+length <= len(primes); start++ {
			sum := prefix[start+length] - prefix[start]
			if sum >= n {
				break
			}
			if isPrime[sum] {
				if length > bestLen {
					bestLen = length
					result = []int{sum}
				} else if length == bestLen {
					result = append(result, sum)
				}
			}
		}
	}

	return result
}