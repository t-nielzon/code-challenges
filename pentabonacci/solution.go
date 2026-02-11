package kata

func CountOddPentaFib(n int) int {
	if n < 0 {
		return 0
	}

	// generate pentabonacci terms up to index n
	f := make([]int64, max(5, n+1))
	f[0] = 0
	f[1] = 1
	f[2] = 1
	f[3] = 2
	f[4] = 4
	for i := 5; i <= n; i++ {
		f[i] = f[i-1] + f[i-2] + f[i-3] + f[i-4] + f[i-5]
	}

	// count distinct odd values; 1 appears twice but should count once
	seen := make(map[int64]bool)
	count := 0
	for i := 0; i <= n; i++ {
		if f[i]%2 != 0 && !seen[f[i]] {
			seen[f[i]] = true
			count++
		}
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}