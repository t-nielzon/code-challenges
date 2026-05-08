package kata

func reverseInt(n int) int {
	r := 0
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	return r
}

func distinctPrimeFactors(n int) []int {
	var factors []int
	for p := 2; p*p <= n; p++ {
		if n%p == 0 {
			factors = append(factors, p)
			for n%p == 0 {
				n /= p
			}
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func sameFactors(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func SameFactRev(nMax int) []int {
	result := []int{}
	for n := 2; n < nMax; n++ {
		rev := reverseInt(n)
		if rev == n {
			continue
		}
		if sameFactors(distinctPrimeFactors(n), distinctPrimeFactors(rev)) {
			result = append(result, n)
		}
	}
	return result
}