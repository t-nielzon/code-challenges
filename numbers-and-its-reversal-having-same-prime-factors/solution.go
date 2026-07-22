package main

import "sort"

func SameFactRev(nMax int) []int {
	result := []int{}
	
	for n := 1; n < nMax; n++ {
		if isPalindrome(n) {
			continue
		}
		
		if sameFactors(n, reverseNumber(n)) {
			result = append(result, n)
		}
	}
	
	sort.Ints(result)
	return result
}

func getPrimeFactors(n int) map[int]bool {
	factors := make(map[int]bool)
	
	if n%2 == 0 {
		factors[2] = true
		for n%2 == 0 {
			n /= 2
		}
	}
	
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			factors[i] = true
			for n%i == 0 {
				n /= i
			}
		}
	}
	
	if n > 1 {
		factors[n] = true
	}
	
	return factors
}

func isPalindrome(n int) bool {
	original := n
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return original == reversed
}

func reverseNumber(n int) int {
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}

func sameFactors(n1, n2 int) bool {
	factors1 := getPrimeFactors(n1)
	factors2 := getPrimeFactors(n2)
	
	if len(factors1) != len(factors2) {
		return false
	}
	
	for k := range factors1 {
		if !factors2[k] {
			return false
		}
	}
	
	return true
}