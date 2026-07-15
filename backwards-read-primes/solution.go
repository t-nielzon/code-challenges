package main

func BackwardsReadPrimes(start, end int) []int {
	result := []int{}

	for num := start; num <= end; num++ {
		if isPrime(num) {
			reversed := reverse(num)
			if isPrime(reversed) && num != reversed {
				result = append(result, num)
			}
		}
	}

	return result
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func reverse(n int) int {
	result := 0
	for n > 0 {
		result = result*10 + n%10
		n /= 10
	}
	return result
}