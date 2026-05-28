package kata

import "strconv"

func MysteryRange(s string, n int) [2]int {
	var target [10]int
	for _, c := range s {
		target[c-'0']++
	}
	for start := 1; start < 100; start++ {
		var cnt [10]int
		for i := start; i < start+n; i++ {
			for _, c := range strconv.Itoa(i) {
				cnt[c-'0']++
			}
		}
		if cnt == target {
			return [2]int{start, start + n - 1}
		}
	}
	return [2]int{0, 0}
}