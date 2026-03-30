package kata

import (
	"strconv"
)

func Smallest(n int64) [3]int64 {
	s := strconv.FormatInt(n, 10)
	bestNum := n
	bestI := int64(0)
	bestJ := int64(0)

	for i := 0; i < len(s); i++ {
		// remove digit at index i
		removed := s[:i] + s[i+1:]
		d := string(s[i])
		for j := 0; j <= len(removed); j++ {
			// insert digit at index j
			candidate := removed[:j] + d + removed[j:]
			num, _ := strconv.ParseInt(candidate, 10, 64)
			if num < bestNum || (num == bestNum && int64(i) < bestI) || (num == bestNum && int64(i) == bestI && int64(j) < bestJ) {
				bestNum = num
				bestI = int64(i)
				bestJ = int64(j)
			}
		}
	}

	return [3]int64{bestNum, bestI, bestJ}
}