package kata

import "math/big"

// the side lengths are fibonacci numbers, so the total perimeter is
// 4 times the sum of the first n+1 fibonacci numbers. that sum equals
// fib(n+2) - 1, giving 4 * (fib(n+2) - 1). big.Int is used since n can
// be large enough to overflow native integer types.
func Perimeter(n *big.Int) *big.Int {
	count := new(big.Int).Add(n, big.NewInt(2)) // iterate up to fib(n+2)

	a := big.NewInt(0) // fib(0)
	b := big.NewInt(1) // fib(1)
	sum := big.NewInt(0)

	for i := big.NewInt(0); i.Cmp(count) < 0; i.Add(i, big.NewInt(1)) {
		sum.Add(sum, a)
		a, b = b, new(big.Int).Add(a, b)
	}

	// sum now holds fib(0)+...+fib(n+1) = fib(n+2)-1
	return sum.Mul(sum, big.NewInt(4))
}