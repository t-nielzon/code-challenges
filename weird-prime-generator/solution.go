package kata

func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gSequence(n int) []uint64 {
	g := make([]uint64, n+1)
	g[0] = 0
	if n >= 1 {
		g[1] = 1
	}
	var prev uint64 = 7
	for i := 2; i <= n; i++ {
		d := gcd(uint64(i), prev)
		cur := prev + d
		g[i] = cur - prev
		prev = cur
	}
	return g
}

func countOnes(n int) int {
	g := gSequence(n)
	c := 0
	for i := 1; i <= n; i++ {
		if g[i] == 1 {
			c++
		}
	}
	return c
}

func maxPn(n int) int {
	seen := map[uint64]bool{}
	var primes []uint64
	var prev uint64 = 7
	for i := 2; len(primes) < n; i++ {
		d := gcd(uint64(i), prev)
		cur := prev + d
		diff := cur - prev
		if diff != 1 {
			if !seen[diff] {
				seen[diff] = true
				primes = append(primes, diff)
			}
		}
		prev = cur
	}
	var mx uint64 = 0
	for _, p := range primes {
		if p > mx {
			mx = p
		}
	}
	return int(mx)
}

func anOverAverage(n int) int {
	var sum uint64 = 0
	count := 0
	var prev uint64 = 7
	for i := 2; count < n; i++ {
		d := gcd(uint64(i), prev)
		cur := prev + d
		if cur-prev != 1 {
			sum += cur / uint64(i)
			count++
		}
		prev = cur
	}
	return int(sum / uint64(n))
}