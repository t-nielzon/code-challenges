// solution.go
package main

import "sort"

func sameFactRev(nMax int) []int {
	result := []int{}
	
	for n := 2; n < nMax; n++ {
		if isPalindrome(n) {
			continue
		}
		
		factors := getPrimeFactors(n)
		reversed := reverse(n)
		reverseFactors := getPrimeFactors(reversed)
		
		if factorsEqual(factors, reverseFactors) {
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

func reverse(n int) int {
	result := 0
	for n > 0 {
		result = result*10 + n%10
		n /= 10
	}
	return result
}

func isPalindrome(n int) bool {
	return n == reverse(n)
}

func factorsEqual(f1, f2 map[int]bool) bool {
	if len(f1) != len(f2) {
		return false
	}
	for key := range f1 {
		if !f2[key] {
			return false
		}
	}
	return true
}