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
	count := 0
	prevIsK := countPrimeFactors(arr[0]) == k
	for i := 1; i < len(arr); i++ {
		currIsK := countPrimeFactors(arr[i]) == k
		if prevIsK && currIsK {
			count++
		}
		prevIsK = currIsK
	}
	return count
}