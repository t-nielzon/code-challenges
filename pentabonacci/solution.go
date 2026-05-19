package kata

import "math/big"

func CountOddPentaFib(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	a := big.NewInt(0)
	b := big.NewInt(1)
	c := big.NewInt(1)
	d := big.NewInt(2)
	e := big.NewInt(4)

	count := 1
	tmp := new(big.Int)

	for i := 5; i <= n; i++ {
		tmp.Add(a, b)
		tmp.Add(tmp, c)
		tmp.Add(tmp, d)
		tmp.Add(tmp, e)
		a, b, c, d, e = b, c, d, e, tmp
		tmp = new(big.Int)
		if e.Bit(0) == 1 {
			count++
		}
	}

	return count
}