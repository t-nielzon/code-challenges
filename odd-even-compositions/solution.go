package main

import "math/big"

// reading a valid composition of N left to right, every part k satisfies
// k == 1, k == running value, or k shares parity with the running value.
// since gray/green parts share parity with the running value (and 1 lowers
// it by exactly one), the remainder after any non-1 part is always even.
// this collapses the recurrence to:
//   f(i) = f(i-1) + sum of f(j) for even j <= i-2
// where f(i-1) accounts for taking a leading 1, and the even-indexed sum
// accounts for every valid same-parity leading part.
func OddEvenCompositions(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	f := make([]*big.Int, n+1)
	f[0] = big.NewInt(1)
	evenSum := big.NewInt(0)
	for i := 1; i <= n; i++ {
		// fold in f(i-2) once it becomes an eligible even index
		if i-2 >= 0 && (i-2)%2 == 0 {
			evenSum = new(big.Int).Add(evenSum, f[i-2])
		}
		f[i] = new(big.Int).Add(f[i-1], evenSum)
	}
	return f[n]
}