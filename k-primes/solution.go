package main

func countPrimeFactors(n int) int {
	if n <= 1 {
		return 0
	}

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
	for n := start; n <= end; n++ {
		if countPrimeFactors(n) == k {
			result = append(result, n)
		}
	}
	return result
}

func puzzle(s int) int {
	primes1 := countKprimes(1, 2, s)
	primes3 := countKprimes(3, 2, s)
	primes7 := countKprimes(7, 2, s)

	set3 := make(map[int]bool)
	for _, p := range primes3 {
		set3[p] = true
	}

	set7 := make(map[int]bool)
	for _, p := range primes7 {
		set7[p] = true
	}

	count := 0

	for _, a := range primes1 {
		if a >= s {
			break
		}
		for _, b := range primes3 {
			if a+b >= s {
				break
			}
			c := s - a - b
			if c > 0 && set7[c] {
				count++
			}
		}
	}

	return count
}