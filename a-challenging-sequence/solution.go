package main

import "sort"

func hashRadSeq(nMax int, k int) int {
	// Calculate radical for each number
	radicals := make([]int, nMax+1)
	for n := 1; n <= nMax; n++ {
		radicals[n] = calculateRadical(n)
	}
	
	// Create pairs of (radical, n)
	type pair struct {
		radical int
		n       int
	}
	pairs := make([]pair, nMax)
	for n := 1; n <= nMax; n++ {
		pairs[n-1] = pair{radicals[n], n}
	}
	
	// Sort by radical, then by n
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].radical != pairs[j].radical {
			return pairs[i].radical < pairs[j].radical
		}
		return pairs[i].n < pairs[j].n
	})
	
	// Return the n value at position k-1 (convert from 1-indexed to 0-indexed)
	return pairs[k-1].n
}

func calculateRadical(n int) int {
	if n == 1 {
		return 1
	}
	
	radical := 1
	temp := n
	
	// Check for factor 2
	if temp%2 == 0 {
		radical *= 2
		for temp%2 == 0 {
			temp /= 2
		}
	}
	
	// Check for odd factors from 3 onwards
	for i := 3; i*i <= temp; i += 2 {
		if temp%i == 0 {
			radical *= i
			for temp%i == 0 {
				temp /= i
			}
		}
	}
	
	// If temp is a prime number greater than 1
	if temp > 1 {
		radical *= temp
	}
	
	return radical
}