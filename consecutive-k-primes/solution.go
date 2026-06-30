package kata

func countPrimeFactors(n int) int {
	count := 0
	for d := 2; d*d <= n; d++ {
		for n%d == 0 {
			count++
			n /= d
		}
	}
	if n > 1 {
		count++
	}
	return count
}

func ConsecKprimes(k int, arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	result := 0
	prev := countPrimeFactors(arr[0])
	for i := 1; i < len(arr); i++ {
		cur := countPrimeFactors(arr[i])
		if prev == k && cur == k {
			result++
		}
		prev = cur
	}
	return result
}