package kprimes

// countPrimeFactors returns the number of prime factors of n counted with
// multiplicity.
func countPrimeFactors(n int) int {
	count := 0
	for d := 2; d*d <= n; d++ {
		for n%d == 0 {
			n /= d
			count++
		}
	}
	if n > 1 {
		count++
	}
	return count
}

func countKprimes(k, start, nd int) []int {
	var res []int
	for n := start; n <= nd; n++ {
		if countPrimeFactors(n) == k {
			res = append(res, n)
		}
	}
	return res
}

func puzzle(s int) int {
	// kOf[i] holds the number of prime factors of i (with multiplicity),
	// so membership in each k-prime set is a single lookup.
	kOf := make([]int, s+1)
	for i := 2; i <= s; i++ {
		if kOf[i] == 0 { // i is prime
			for j := i; j <= s; j += i {
				m := j
				for m%i == 0 {
					m /= i
					kOf[j]++
				}
			}
		}
	}

	total := 0
	for a := 2; a <= s; a++ {
		if kOf[a] != 1 {
			continue
		}
		for b := 2; a+b <= s; b++ {
			if kOf[b] != 3 {
				continue
			}
			c := s - a - b
			if c >= 2 && kOf[c] == 7 {
				total++
			}
		}
	}
	return total
}