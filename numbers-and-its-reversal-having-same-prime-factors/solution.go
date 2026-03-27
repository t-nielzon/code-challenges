package kata

func SameFactRev(nMax int) []int {
	var result []int
	for n := 10; n < nMax; n++ {
		rev := reverseNum(n)
		if rev == n || rev >= nMax {
			continue
		}
		if samePrimeFactors(n, rev) {
			result = append(result, n)
		}
	}
	return result
}

func reverseNum(n int) int {
	rev := 0
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev
}

func primeFactors(n int) map[int]bool {
	factors := make(map[int]bool)
	for d := 2; d*d <= n; d++ {
		for n%d == 0 {
			factors[d] = true
			n /= d
		}
	}
	if n > 1 {
		factors[n] = true
	}
	return factors
}

func samePrimeFactors(a, b int) bool {
	fa := primeFactors(a)
	fb := primeFactors(b)
	if len(fa) != len(fb) {
		return false
	}
	for k := range fa {
		if !fb[k] {
			return false
		}
	}
	return true
}