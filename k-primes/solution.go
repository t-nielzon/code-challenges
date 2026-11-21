package main

func countPrimeFactors(n int) int {
	count := 0
	for n%2 == 0 {
		count++
		n /= 2
	}
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			count++
			n /= i
		}
	}
	if n > 1 {
		count++
	}
	return count
}

func countKprimes(k, start, end int) []int {
	var result []int
	for i := start; i <= end; i++ {
		if countPrimeFactors(i) == k {
			result = append(result, i)
		}
	}
	return result
}

func puzzle(s int) int {
	primes1 := make(map[int]bool)
	primes3 := make(map[int]bool)
	primes7 := make(map[int]bool)
	
	for i := 2; i < s; i++ {
		factors := countPrimeFactors(i)
		if factors == 1 {
			primes1[i] = true
		}
		if factors == 3 {
			primes3[i] = true
		}
		if factors == 7 {
			primes7[i] = true
		}
	}
	
	count := 0
	for a := range primes1 {
		for b := range primes3 {
			c := s - a - b
			if c > 0 && primes7[c] {
				count++
			}
		}
	}
	
	return count
}