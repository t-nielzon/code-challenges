package kata

func CountOddPentaFib(n int) int {
	if n < 0 {
		return 0
	}

	// base cases for the pentabonacci sequence
	base := [5]int{0, 1, 1, 2, 4}

	// for small n, count odd values among base[0..n]
	if n < 5 {
		count := 0
		seenOne := false
		for i := 0; i <= n; i++ {
			if base[i]%2 != 0 {
				// 1 appears at index 1 and 2, count it only once
				if base[i] == 1 {
					if !seenOne {
						seenOne = true
						count++
					}
				} else {
					count++
				}
			}
		}
		return count
	}

	// the parity pattern of pentabonacci repeats with period 12:
	// indices 0-11: even, odd, odd, even, even, even, even, odd, odd, even, even, even
	// after the initial pair of 1s, no further duplicates occur among odd values
	// 1 is counted once, so we start with count=1 having already handled indices 0..4
	// in base cases index 1 and 2 are odd (both 1, counted once), indices 3,4 are even
	count := 1 // for the value 1 (indices 1 and 2)

	// odd positions in one full cycle of 12 (0-indexed within cycle): 1,2,7,8
	// but we already counted indices 0..4, so start from index 5
	// from index 5 onward, within each cycle of 12 starting at 0:
	// positions 7,8 are odd in cycle — that's offset 7 and 8
	// then next cycle positions 1,2,7,8 are odd — offsets 13,14,19,20, etc.

	// simpler: just iterate and check parity using the recurrence on parity bits
	w := [5]int{base[0] & 1, base[1] & 1, base[2] & 1, base[3] & 1, base[4] & 1}

	for i := 5; i <= n; i++ {
		p := (w[0] + w[1] + w[2] + w[3] + w[4]) & 1
		if p == 1 {
			count++
		}
		// shift window
		w[0] = w[1]
		w[1] = w[2]
		w[2] = w[3]
		w[3] = w[4]
		w[4] = p
	}

	return count
}