package kata

import "strconv"

func Values(n int) int {
	seen := make(map[int]bool)

	for i := 1; i*i < n; i++ {
		sum := 0
		for j := i; sum < n; j++ {
			sum += j * j
			if j > i && sum < n && isPalindrome(sum) {
				seen[sum] = true
			}
		}
	}

	return len(seen)
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}