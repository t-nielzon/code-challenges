package kata

import "math/big"

func KthLastDigPrime(k int) []int {
	mod := big.NewInt(1000000000)
	// a term qualifies only when it genuinely has more than nine digits,
	// so we require the value to be at least 10^9 before taking its tail.
	threshold := big.NewInt(1000000000)

	// seeds f(0)..f(4)
	f := []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(4),
	}

	last9 := new(big.Int)
	count := 0
	n := 4

	for {
		val := f[len(f)-1]
		if val.Cmp(threshold) >= 0 {
			last9.Mod(val, mod)
			if last9.ProbablyPrime(20) {
				count++
				if count == k {
					return []int{n, int(last9.Int64())}
				}
			}
		}

		// next = f(n-1) + f(n-2) - f(n-3) + f(n-4) - f(n-5)
		next := new(big.Int).Set(f[len(f)-1])
		next.Add(next, f[len(f)-2])
		next.Sub(next, f[len(f)-3])
		next.Add(next, f[len(f)-4])
		next.Sub(next, f[len(f)-5])
		f = append(f, next)

		// only the last five values feed the recurrence, keep the slice bounded
		if len(f) > 6 {
			f = f[len(f)-6:]
		}
		n++
	}
}