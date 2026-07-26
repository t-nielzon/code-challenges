package main

func countPrimeFactors(n int) int {
	if n <= 1 {
		return 0
	}
	count := 0
	// divide out all 2s
	for n%2 == 0 {
		count++
		n /= 2
	}
	// check odd factors from 3 onwards
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			count++
			n /= i
		}
	}
	// if n > 1, then it's a prime factor itself
	if n > 1 {
		count++
	}
	return count
}

func kprimes_step(k, step, start, nd int) [][]int {
	var result [][]int
	
	for num := start; num <= nd; num++ {
		if countPrimeFactors(num) == k {
			nextNum := num + step
			if nextNum <= nd && countPrimeFactors(nextNum) == k {
				result = append(result, []int{num, nextNum})
			}
		}
	}
	
	return result
}