package kata

func KprimesStep(k, step, start, nd uint64) [][]uint64 {
	result := [][]uint64{}
	countFactors := func(n uint64) uint64 {
		var c uint64
		for p := uint64(2); p*p <= n; p++ {
			for n%p == 0 {
				n /= p
				c++
			}
		}
		if n > 1 {
			c++
		}
		return c
	}
	var prev uint64
	hasPrev := false
	for n := start; n <= nd; n++ {
		if n < 2 {
			continue
		}
		if countFactors(n) == k {
			if hasPrev && n-prev == step {
				result = append(result, []uint64{prev, n})
			}
			prev = n
			hasPrev = true
		}
	}
	return result
}