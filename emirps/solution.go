package main

func FindEmirp(n int) [3]int {
	if n <= 2 {
		return [3]int{0, 0, 0}
	}
	
	// Generate all primes below n using Sieve of Eratosthenes
	primes := sieve(n)
	
	// Create a map for O(1) prime lookup
	primeSet := make(map[int]bool)
	for _, p := range primes {
		primeSet[p] = true
	}
	
	count := 0
	largest := 0
	sum := 0
	
	// Check each prime to see if it's an emirp
	for _, p := range primes {
		reversed := reverse(p)
		// emirp: reversed is a different prime
		if reversed != p && primeSet[reversed] {
			count++
			if p > largest {
				largest = p
			}
			sum += p
		}
	}
	
	return [3]int{count, largest, sum}
}

func sieve(n int) []int {
	if n < 2 {
		return []int{}
	}
	
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
	
	var primes []int
	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	
	return primes
}

func reverse(n int) int {
	result := 0
	for n > 0 {
		result = result*10 + n%10
		n /= 10
	}
	return result
}