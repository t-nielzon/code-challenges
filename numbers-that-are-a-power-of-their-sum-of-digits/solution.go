// solution.go
package kata

import "sort"

func digitSum(n int) int {
	s := 0
	for n > 0 {
		s += n % 10
		n /= 10
	}
	return s
}

func PowerSumDigTerm(n int) int {
	seen := map[int]bool{}

	// iterate over possible digit sums and exponents
	for base := 2; base <= 200; base++ {
		power := base * base // start at exponent 2
		for power > 0 && power < 1e12 {
			if power >= 10 && digitSum(power) == base {
				seen[power] = true
			}
			power *= base
			if power/base != power/base { // overflow guard
				break
			}
		}
	}

	results := make([]int, 0, len(seen))
	for v := range seen {
		results = append(results, v)
	}
	sort.Ints(results)

	return results[n-1]
}