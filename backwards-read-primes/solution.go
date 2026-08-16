package main

import "strconv"

func BackwardsReadPrimes(start, end int) []int {
	result := []int{}

	for num := start; num <= end; num++ {
		if isPrime(num) {
			reversed := reverseNumber(num)
			if reversed != num && isPrime(reversed) {
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

func reverseNumber(n int) int {
	str := strconv.Itoa(n)
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	num, _ := strconv.Atoi(reversed)
	return num
}