package kata

import "math/big"

func CountOddPentaFib(n int) int {
	if n < 0 {
		return 0
	}

	a := [5]*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(4),
	}

	one := big.NewInt(1)
	two := big.NewInt(2)
	tmp := new(big.Int)
	count := 0
	seenOne := false

	for i := 0; i <= n; i++ {
		var val *big.Int
		if i < 5 {
			val = a[i]
		} else {
			next := new(big.Int)
			for j := 0; j < 5; j++ {
				next.Add(next, a[j])
			}
			copy(a[:], a[1:])
			a[4] = next
			val = next
		}

		tmp.And(val, one)
		if tmp.Sign() != 0 {
			if val.Cmp(one) == 0 {
				if !seenOne {
					seenOne = true
					count++
				}
			} else {
				count++
			}
		}
		_ = two
	}

	return count
}