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

func KprimesStep(k, step, start, nd int) [][]int {
	result := [][]int{}
	var prev = -1
	for n := start; n <= nd; n++ {
		if countPrimeFactors(n) == k {
			if prev != -1 && n-prev == step {
				result = append(result, []int{prev, n})
			}
			prev = n
		}
	}
	return result
}