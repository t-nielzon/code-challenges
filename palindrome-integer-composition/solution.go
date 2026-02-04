package kata

func Values(n int) int {
	palindromes := make(map[int]bool)

	// Find all sums of consecutive squares less than n
	// Start from each positive integer i
	for i := 1; i*i < n; i++ {
		sum := 0
		// Add consecutive squares starting from i
		for j := i; sum < n; j++ {
			sum += j * j
			// We need at least 2 consecutive squares (j > i means we have i and at least one more)
			if j > i && sum < n && isPalindrome(sum) {
				palindromes[sum] = true
			}
		}
	}

	return len(palindromes)
}

func isPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	original := n
	reversed := 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return original == reversed
}