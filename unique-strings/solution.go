package kata

import (
	"math/big"
	"strings"
)

func UniqCount(s string) *big.Int {
	s = strings.ToLower(s)

	freq := make(map[rune]int)
	for _, c := range s {
		freq[c]++
	}

	result := factorial(len(s))
	for _, count := range freq {
		result.Div(result, factorial(count))
	}

	return result
}

func factorial(n int) *big.Int {
	r := big.NewInt(1)
	for i := 2; i <= n; i++ {
		r.Mul(r, big.NewInt(int64(i)))
	}
	return r
}