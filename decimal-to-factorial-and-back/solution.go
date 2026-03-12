package kata

import "math/big"

func DecToFact(nb int64) string {
	n := big.NewInt(nb)
	zero := big.NewInt(0)
	digits := []byte{}
	base := big.NewInt(1)

	for i := int64(1); n.Cmp(zero) > 0; i++ {
		base.SetInt64(i)
		mod := new(big.Int)
		n.DivMod(n, base, mod)
		d := mod.Int64()
		if d < 10 {
			digits = append(digits, byte('0'+d))
		} else {
			digits = append(digits, byte('A'+d-10))
		}
	}

	if len(digits) == 0 {
		return "0"
	}

	// reverse
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}

func FactToDec(s string) int64 {
	result := big.NewInt(0)
	fact := big.NewInt(1)

	n := len(s)
	for i := 1; i < n; i++ {
		fact.Mul(fact, big.NewInt(int64(i)))
	}

	for i := 0; i < n; i++ {
		c := s[i]
		var d int64
		if c >= '0' && c <= '9' {
			d = int64(c - '0')
		} else {
			d = int64(c-'A') + 10
		}
		term := new(big.Int).Mul(big.NewInt(d), fact)
		result.Add(result, term)

		base := int64(n - 1 - i)
		if base > 0 {
			fact.Div(fact, big.NewInt(base))
		}
	}

	return result.Int64()
}