package main

import (
	"sort"
	"strconv"
)

func reverse(n int) int {
	s := strconv.Itoa(n)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversed, _ := strconv.Atoi(string(runes))
	return reversed
}

func isPalindrome(n int) bool {
	return n == reverse(n)
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

func SameFactRev(nMax int) []int {
	result := []int{}
	
	for n := 2; n < nMax; n++ {
		if isPalindrome(n) {
			continue
		}
		
		factorsN := getPrimeFactors(n)
		factorsRev := getPrimeFactors(reverse(n))
		
		if len(factorsN) == len(factorsRev) {
			same := true
			for k := range factorsN {
				if !factorsRev[k] {
					same = false
					break
				}
			}
			if same {
				result = append(result, n)
			}
		}
	}
	
	sort.Ints(result)
	return result
}