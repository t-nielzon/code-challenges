package kata

import "math/big"

func KthLastDigPrime(k int) [2]int {
	mod := big.NewInt(1000000000)
	threshold := big.NewInt(1000000000)

	v0 := big.NewInt(0)
	v1 := big.NewInt(1)
	v2 := big.NewInt(1)
	v3 := big.NewInt(2)
	v4 := big.NewInt(4)

	count := 0
	n := 4
	for {
		n++
		next := new(big.Int).Set(v4)
		next.Add(next, v3)
		next.Sub(next, v2)
		next.Add(next, v1)
		next.Sub(next, v0)

		if next.Cmp(threshold) >= 0 {
			last9 := new(big.Int).Mod(next, mod)
			if last9.ProbablyPrime(20) {
				count++
				if count == k {
					return [2]int{n, int(last9.Int64())}
				}
			}
		}

		v0, v1, v2, v3, v4 = v1, v2, v3, v4, next
	}
}