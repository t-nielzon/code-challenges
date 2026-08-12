package kata

func OddEvenCompositions(n int) int {
	memo := make(map[int]int)
	return count(n, memo)
}

func count(n int, memo map[int]int) int {
	if n == 0 {
		return 1
	}

	if val, exists := memo[n]; exists {
		return val
	}

	// we can always choose 1, or choose n itself
	result := count(n-1, memo) + 1

	// choose k where k has same parity as n and 2 <= k < n
	start := 2
	if n%2 == 1 {
		start = 3
	}

	for k := start; k < n; k += 2 {
		result += count(n-k, memo)
	}

	memo[n] = result
	return result
}