package main

import "strconv"

func values(n int) int {
	isPalindrome := func(num int) bool {
		s := strconv.Itoa(num)
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	palindromes := make(map[int]bool)

	// for each starting position of consecutive squares
	for i := 1; i*i < n; i++ {
		sum := i * i
		// for each ending position (must be at least 2 consecutive squares)
		for j := i + 1; sum < n; j++ {
			sum += j * j
			if sum >= n {
				break
			}
			if isPalindrome(sum) {
				palindromes[sum] = true
			}
		}
	}

	return len(palindromes)
}