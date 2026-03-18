package kata

import "math/big"

func Perimeter(n *big.Int) *big.Int {
	limit := n.Int64()
	a := big.NewInt(1)
	b := big.NewInt(1)
	sum := big.NewInt(1)

	for i := int64(1); i <= limit; i++ {
		a, b = b, new(big.Int).Add(a, b)
		sum.Add(sum, b)
	}

	return sum.Mul(sum, big.NewInt(4))
}