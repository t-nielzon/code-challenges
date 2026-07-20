func LastFibDigit(n int64) int64 {
	// the last digits of Fibonacci numbers repeat with period 60 (Pisano period for modulo 10)
	// map n to its equivalent position in the cycle
	n = ((n - 1) % 60) + 1
	
	if n == 1 || n == 2 {
		return 1
	}
	
	prev, curr := int64(1), int64(1)
	for i := int64(3); i <= n; i++ {
		prev, curr = curr, (prev+curr)%10
	}
	
	return curr
}