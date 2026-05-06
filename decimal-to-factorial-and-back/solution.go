package kata

import (
	"math/big"
	"strings"
)

const factAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Dec2FactString(nb *big.Int) string {
	n := new(big.Int).Set(nb)
	if n.Sign() == 0 {
		return "0"
	}
	var digits []byte
	divisor := big.NewInt(1)
	one := big.NewInt(1)
	rem := new(big.Int)
	for n.Sign() > 0 {
		n.DivMod(n, divisor, rem)
		digits = append(digits, factAlphabet[rem.Int64()])
		divisor = new(big.Int).Add(divisor, one)
	}
	var sb strings.Builder
	for i := len(digits) - 1; i >= 0; i-- {
		sb.WriteByte(digits[i])
	}
	return sb.String()
}

func FactString2Dec(s string) *big.Int {
	result := big.NewInt(0)
	fact := big.NewInt(1)
	for p := 0; p < len(s); p++ {
		c := s[len(s)-1-p]
		var d int64
		switch {
		case c >= '0' && c <= '9':
			d = int64(c - '0')
		case c >= 'A' && c <= 'Z':
			d = int64(c-'A') + 10
		}
		term := new(big.Int).Mul(fact, big.NewInt(d))
		result.Add(result, term)
		fact = new(big.Int).Mul(fact, big.NewInt(int64(p+1)))
	}
	return result
}