package kata

import "math/big"

// brahmagupta–fibonacci identity:
// (a^2+b^2)(c^2+d^2) = (ac+bd)^2 + (ad-bc)^2
// fold the pairs left to right keeping a single (A, B) so the running
// product P always stays expressed as A^2 + B^2, regardless of magnitude.
func Solve(arr []int) [2]*big.Int {
	a := big.NewInt(int64(arr[0]))
	b := big.NewInt(int64(arr[1]))

	tmp := new(big.Int)
	for i := 2; i < len(arr); i += 2 {
		c := big.NewInt(int64(arr[i]))
		d := big.NewInt(int64(arr[i+1]))

		// newA = a*c + b*d
		newA := new(big.Int).Mul(a, c)
		newA.Add(newA, tmp.Mul(b, d))

		// newB = a*d - b*c (absolute value, since sign is irrelevant for (1))
		newB := new(big.Int).Mul(a, d)
		newB.Sub(newB, tmp.Mul(b, c))
		newB.Abs(newB)

		a, b = newA, newB
	}

	return [2]*big.Int{a, b}
}