package kata

func CountOddPentaFib(n int) int {
	if n < 0 {
		return 0
	}

	// generate pentabonacci terms up to index n
	f := make([]uint64, max(5, n+1))
	f[0] = 0
	f[1] = 1
	f[2] = 1
	f[3] = 2
	f[4] = 4
	for i := 5; i <= n; i++ {
		f[i] = f[i-1] + f[i-2] + f[i-3] + f[i-4] + f[i-5]
	}

	count := 0
	seenOne := false
	for i := 0; i <= n; i++ {
		if f[i]%2 == 1 {
			if f[i] == 1 {
				if !seenOne {
					seenOne = true
					count++
				}
			} else {
				count++
			}
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