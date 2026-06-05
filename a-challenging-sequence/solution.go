package kata

import "sort"

// rad uses a smallest-prime-factor sieve so radicals are built in a single
// pass, keeping it fast for the large nMax the tests use.
func hashRadSeq(nMax int, k int) int {
	rad := make([]int, nMax+1)
	for i := 1; i <= nMax; i++ {
		rad[i] = 1
	}

	// for every prime p, multiply its radical contribution into all multiples
	for p := 2; p <= nMax; p++ {
		if rad[p] == 1 { // p is prime (no smaller factor touched it yet)
			for m := p; m <= nMax; m += p {
				rad[m] *= p
			}
		}
	}

	nums := make([]int, nMax)
	for i := 1; i <= nMax; i++ {
		nums[i-1] = i
	}

	sort.Slice(nums, func(a, b int) bool {
		na, nb := nums[a], nums[b]
		if rad[na] != rad[nb] {
			return rad[na] < rad[nb]
		}
		return na < nb
	})

	return nums[k-1]
}