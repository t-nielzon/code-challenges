package kata

func FindPrimesSextuplet(sumLimit int) [6]int {
	offsets := [6]int{0, 4, 6, 10, 12, 16}

	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		if n < 4 {
			return true
		}
		if n%2 == 0 || n%3 == 0 {
			return false
		}
		for i := 5; i*i <= n; i += 6 {
			if n%i == 0 || n%(i+2) == 0 {
				return false
			}
		}
		return true
	}

	for p := 7; ; p += 2 {
		if !isPrime(p) {
			continue
		}
		valid := true
		for i := 1; i < 6; i++ {
			if !isPrime(p + offsets[i]) {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		sum := 6*p + 48
		if sum > sumLimit {
			var result [6]int
			for i, o := range offsets {
				result[i] = p + o
			}
			return result
		}
	}
}