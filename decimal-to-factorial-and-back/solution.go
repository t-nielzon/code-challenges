package kata

import "math/big"

const digits = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Dec2FactString(nb int64) string {
	n := big.NewInt(nb)
	i := 1
	var result []byte
	for n.Sign() > 0 {
		mod := new(big.Int)
		n.DivMod(n, big.NewInt(int64(i)), mod)
		result = append(result, digits[int(mod.Int64())])
		i++
	}
	if len(result) == 0 {
		return "0"
	}
	// reverse
	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}
	return string(result)
}

func FactString2Dec(s string) int64 {
	result := big.NewInt(0)
	n := len(s) - 1
	fact := big.NewInt(1)
	for i := 1; i <= n; i++ {
		fact.Mul(fact, big.NewInt(int64(i)))
	}
	for i, c := range s {
		var val int64
		if c >= '0' && c <= '9' {
			val = int64(c - '0')
		} else {
			val = int64(c-'A') + 10
		}
		term := new(big.Int).Mul(big.NewInt(val), fact)
		result.Add(result, term)
		base := int64(n - i)
		if base > 0 {
			fact.Div(fact, big.NewInt(base))
		}
	}
	return result.Int64()
}