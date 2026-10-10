// solution.go
package main

import "strconv"

func SameFactRev(nMax int) []int {
	result := []int{}
	
	for num := 2; num < nMax; num++ {
		if isPalindrome(num) {
			continue
		}
		
		factors1 := getPrimeFactors(num)
		factors2 := getPrimeFactors(reverseNumber(num))
		
		if haveSamePrimeFactors(factors1, factors2) {
			result = append(result, num)
		}
	}
	
	return result
}

func isPalindrome(num int) bool {
	s := strconv.Itoa(num)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func reverseNumber(num int) int {
	rev := 0
	for num > 0 {
		rev = rev*10 + num%10
		num /= 10
	}
	return rev
}

func getPrimeFactors(num int) map[int]bool {
	factors := make(map[int]bool)
	
	if num%2 == 0 {
		factors[2] = true
		for num%2 == 0 {
			num /= 2
		}
	}
	
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			factors[i] = true
			for num%i == 0 {
				num /= i
			}
		}
	}
	
	if num > 1 {
		factors[num] = true
	}
	
	return factors
}

func haveSamePrimeFactors(f1, f2 map[int]bool) bool {
	if len(f1) != len(f2) {
		return false
	}
	for k := range f1 {
		if !f2[k] {
			return false
		}
	}
	return true
}