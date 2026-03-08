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

func CountKprimes(k, start, nd int) []int {
	var result []int
	for n := start; n <= nd; n++ {
		if countPrimeFactors(n) == k {
			result = append(result, n)
		}
	}
	return result
}

func Puzzle(s int) int {
	ones := CountKprimes(1, 2, s)
	threes := CountKprimes(3, 2, s)
	sevens := CountKprimes(7, 2, s)
	count := 0
	for _, c := range sevens {
		for _, b := range threes {
			if b+c >= s {
				break
			}
			rem := s - b - c
			if rem < 2 {
				break
			}
			if countPrimeFactors(rem) == 1 {
				count++
			}
		}
	}
	_ = ones
	return count
}