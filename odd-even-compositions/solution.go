package kata

import "math/big"

func OddEvenCompositions(n int) *big.Int {
	f := make([]*big.Int, n+1)
	f[0] = big.NewInt(1)
	for i := 1; i <= n; i++ {
		sum := new(big.Int).Set(f[i-1])
		start := 2
		if i%2 == 1 {
			start = 3
		}
		for k := start; k <= i; k += 2 {
			sum.Add(sum, f[i-k])
		}
		f[i] = sum
	}
	return f[n]
}