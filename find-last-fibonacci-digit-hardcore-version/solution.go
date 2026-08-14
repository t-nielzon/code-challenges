package main

func LastFibDigit(n int64) int {
	// the last digits of fibonacci numbers follow the Pisano period: they repeat every 60 numbers
	n = n % 60
	if n == 0 {
		n = 60
	}

	if n == 1 || n == 2 {
		return 1
	}

	prev, curr := 1, 1
	for i := int64(3); i <= n; i++ {
		prev, curr = curr, (prev+curr)%10
	}

	return curr
}