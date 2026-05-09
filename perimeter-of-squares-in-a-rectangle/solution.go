package kata

import "math/big"

func Perimeter(n uint64) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)
	sum := big.NewInt(0)
	for i := uint64(0); i <= n; i++ {
		sum.Add(sum, b)
		a.Add(a, b)
		a, b = b, a
	}
	return sum.Mul(sum, big.NewInt(4))
}