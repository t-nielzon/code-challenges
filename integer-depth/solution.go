package kata

import "strconv"

func ComputeDepth(n int) int {
	seen := [10]bool{}
	count := 0
	for i := 1; count < 10; i++ {
		for _, d := range strconv.Itoa(n * i) {
			if !seen[d-'0'] {
				seen[d-'0'] = true
				count++
			}
		}
		if count == 10 {
			return i
		}
	}
	return 0
}