package kata

import "math/big"

func GenerateDiagonal(n, l int) []*big.Int {
	result := make([]*big.Int, l)
	if l == 0 {
		return result
	}
	result[0] = big.NewInt(1)
	for k := 1; k < l; k++ {
		// each element is C(n+k, n), derived from the previous via the
		// recurrence C(n+k, n) = C(n+k-1, n) * (n+k) / k
		val := new(big.Int).Mul(result[k-1], big.NewInt(int64(n+k)))
		val.Div(val, big.NewInt(int64(k)))
		result[k] = val
	}
	return result
}