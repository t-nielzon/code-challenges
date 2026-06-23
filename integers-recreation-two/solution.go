package kata

import "sort"

func prod2sum(a, b, c, d int) [][]int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	// brahmagupta–fibonacci identity gives the two decompositions
	pairs := [][]int{
		{abs(a*c - b*d), abs(a*d + b*c)},
		{abs(a*c + b*d), abs(a*d - b*c)},
	}

	seen := map[[2]int]bool{}
	var res [][]int
	for _, p := range pairs {
		lo, hi := p[0], p[1]
		if lo > hi {
			lo, hi = hi, lo
		}
		key := [2]int{lo, hi}
		if seen[key] {
			continue
		}
		seen[key] = true
		res = append(res, []int{lo, hi})
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i][0] != res[j][0] {
			return res[i][0] < res[j][0]
		}
		return res[i][1] < res[j][1]
	})

	return res
}