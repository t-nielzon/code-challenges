package main

import "sort"

func HashRadSeq(nMax, k int) int {
	// Sieve of Eratosthenes to find all primes up to nMax
	isPrime := make([]bool, nMax+1)
	for i := 2; i <= nMax; i++ {
		isPrime[i] = true
	}
	for p := 2; p*p <= nMax; p++ {
		if isPrime[p] {
			for multiple := p * p; multiple <= nMax; multiple += p {
				isPrime[multiple] = false
			}
		}
	}

	// Calculate radical for each number
	radicals := make([]int, nMax+1)
	for i := 1; i <= nMax; i++ {
		radicals[i] = 1
	}

	for p := 2; p <= nMax; p++ {
		if isPrime[p] {
			for multiple := p; multiple <= nMax; multiple += p {
				radicals[multiple] *= p
			}
		}
	}

	// Create pairs of (n, radical)
	type pair struct {
		n      int
		radical int
	}

	pairs := make([]pair, nMax)
	for i := 1; i <= nMax; i++ {
		pairs[i-1] = pair{n: i, radical: radicals[i]}
	}

	// Sort by radical first, then by n for ties
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].radical != pairs[j].radical {
			return pairs[i].radical < pairs[j].radical
		}
		return pairs[i].n < pairs[j].n
	})

	// Return the n value at position k (k is 1-indexed)
	return pairs[k-1].n
}