package kata

import (
	"math/big"
	"strings"
)

func UniqCount(s string) *big.Int {
	// case insensitive: fold everything to one case before counting
	s = strings.ToLower(s)

	counts := make(map[rune]int)
	n := 0
	for _, r := range s {
		counts[r]++
		n++
	}

	// permutations of a multiset: n! / (c1! * c2! * ... * ck!)
	result := factorial(n)
	for _, c := range counts {
		result.Div(result, factorial(c))
	}
	return result
}

func factorial(n int) *big.Int {
	f := big.NewInt(1)
	for i := 2; i <= n; i++ {
		f.Mul(f, big.NewInt(int64(i)))
	}
	return f
}