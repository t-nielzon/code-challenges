package main

func FindEmirp(n int) [3]int {
	if n <= 2 {
		return [3]int{0, 0, 0}
	}

	primes := sievePrimes(n)
	primeSet := make(map[int]bool)
	for _, p := range primes {
		primeSet[p] = true
	}

	count := 0
	largest := 0
	sum := 0

	for _, prime := range primes {
		reversed := reverseNumber(prime)

		if reversed != prime && primeSet[reversed] {
			count++
			largest = prime
			sum += prime
		}
	}

	return [3]int{count, largest, sum}
}

func sievePrimes(n int) []int {
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

func reverseNumber(num int) int {
	result := 0
	for num > 0 {
		result = result*10 + num%10
		num /= 10
	}
	return result
}