package kata

import "sort"

func PrimesLongestSum(n int) []int {
	result := []int{}
	if n <= 2 {
		return result
	}

	sieve := make([]bool, n)
	for i := 2; i < n; i++ {
		sieve[i] = true
	}
	for i := 2; i*i < n; i++ {
		if sieve[i] {
			for j := i * i; j < n; j += i {
				sieve[j] = false
			}
		}
	}

	primes := []int{}
	for i := 2; i < n; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}

	bestLen := make(map[int]int)
	for i := 0; i < len(primes); i++ {
		sum := 0
		for j := i; j < len(primes); j++ {
			sum += primes[j]
			if sum >= n {
				break
			}
			if sieve[sum] {
				length := j - i + 1
				if length > bestLen[sum] {
					bestLen[sum] = length
				}
			}
		}
	}

	maxLen := 0
	for _, l := range bestLen {
		if l > maxLen {
			maxLen = l
		}
	}

	for p, l := range bestLen {
		if l == maxLen {
			result = append(result, p)
		}
	}
	sort.Ints(result)
	return result
}