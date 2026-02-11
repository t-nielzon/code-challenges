package kata

import (
	"math/big"
)

const mod int64 = 1000000000

func KthLastDigPrime(k int) [2]int64 {
	// keep only last 9 digits to check primality
	f := [5]int64{0, 1, 1, 2, 4}
	n := 4
	count := 0

	// helper to ensure positive mod result
	posMod := func(v, m int64) int64 {
		v %= m
		if v < 0 {
			v += m
		}
		return v
	}

	// need to also track whether the full value has >= 10 digits
	// use a separate big.Int for magnitude tracking
	bf := [5]*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(4),
	}
	threshold := new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil) // 10^9

	for {
		n++
		// indices into circular-style, but simpler to just shift
		next := f[4] + f[3] - f[2] + f[1] - f[0]
		next = posMod(next, mod)

		// track big value for digit count
		bNext := new(big.Int)
		bNext.Add(bf[4], bf[3])
		bNext.Sub(bNext, bf[2])
		bNext.Add(bNext, bf[1])
		bNext.Sub(bNext, bf[0])

		// shift window
		f[0], f[1], f[2], f[3], f[4] = f[1], f[2], f[3], f[4], next
		bf[0], bf[1], bf[2], bf[3], bf[4] = bf[1], bf[2], bf[3], bf[4], bNext

		// only consider terms where the full number has more than 9 digits
		if bNext.Cmp(threshold) >= 0 {
			if big.NewInt(next).ProbablyPrime(20) {
				count++
				if count == k {
					return [2]int64{int64(n), next}
				}
			}
		}
	}
}