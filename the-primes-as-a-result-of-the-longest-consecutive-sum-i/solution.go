package kata

import "sort"

func LongestConsecutive(n int) []int {
	if n < 3 {
		return []int{}
	}

	// sieve of eratosthenes for all numbers below n
	isComposite := make([]bool, n)
	primes := []int{}
	for i := 2; i < n; i++ {
		if !isComposite[i] {
			primes = append(primes, i)
			for j := i * 2; j < n; j += i {
				isComposite[j] = true
			}
		}
	}

	isPrime := func(v int) bool {
		return v >= 2 && v < n && !isComposite[v]
	}

	// for every prime reachable as a consecutive sum, keep its longest chain
	bestLen := map[int]int{}
	maxLen := 0
	for i := 0; i < len(primes); i++ {
		sum := 0
		for j := i; j < len(primes); j++ {
			sum += primes[j]
			if sum >= n {
				break
			}
			if isPrime(sum) {
				length := j - i + 1
				if length > bestLen[sum] {
					bestLen[sum] = length
				}
				if length > maxLen {
					maxLen = length
				}
			}
		}
	}

	result := []int{}
	for p, l := range bestLen {
		if l == maxLen {
			result = append(result, p)
		}
	}
	sort.Ints(result)
	return result
}