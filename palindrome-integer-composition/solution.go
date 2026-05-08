package kata

import "strconv"

func isPalindrome(x uint64) bool {
	s := strconv.FormatUint(x, 10)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func Values(n uint64) int {
	found := make(map[uint64]struct{})
	for i := uint64(1); i*i < n; i++ {
		sum := i * i
		for j := i + 1; ; j++ {
			sum += j * j
			if sum >= n {
				break
			}
			if isPalindrome(sum) {
				found[sum] = struct{}{}
			}
		}
	}
	return len(found)
}