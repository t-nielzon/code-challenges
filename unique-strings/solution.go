package kata

import (
	"math/big"
	"strings"
)

func UniqCount(s string) *big.Int {
	s = strings.ToUpper(s)
	n := len(s)

	freq := make(map[rune]int)
	for _, c := range s {
		freq[c]++
	}

	result := factorial(n)
	for _, count := range freq {
		if count > 1 {
			result.Div(result, factorial(count))
		}
	}

	return result
}

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}