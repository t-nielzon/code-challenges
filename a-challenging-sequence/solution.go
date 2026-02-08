package kata

import "sort"

func HashRadSeq(nMax, k int) int {
	// compute radical for each number 1..nMax using a sieve approach
	rad := make([]int, nMax+1)
	for i := range rad {
		rad[i] = 1
	}
	for p := 2; p <= nMax; p++ {
		if rad[p] == 1 { // p is prime
			for j := p; j <= nMax; j += p {
				rad[j] *= p
			}
		}
	}

	// build list of numbers 1..nMax and sort by (radical, n)
	nums := make([]int, nMax)
	for i := 0; i < nMax; i++ {
		nums[i] = i + 1
	}
	sort.Slice(nums, func(i, j int) bool {
		ri, rj := rad[nums[i]], rad[nums[j]]
		if ri != rj {
			return ri < rj
		}
		return nums[i] < nums[j]
	})

	return nums[k-1]
}