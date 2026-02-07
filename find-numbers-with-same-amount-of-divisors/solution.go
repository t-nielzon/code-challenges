package kata

func CountPairsInt(diff, nMax int) int {
	count := 0
	for prev := 1; prev+diff < nMax; prev++ {
		next := prev + diff
		if countDivisors(prev) == countDivisors(next) {
			count++
		}
	}
	return count
}

func countDivisors(n int) int {
	if n == 1 {
		return 1
	}
	count := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			count++
			if i != n/i {
				count++
			}
		}
	}
	return count
}