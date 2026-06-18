package kata

import "strconv"

func Smallest(n int64) [3]int64 {
	s := []byte(strconv.FormatInt(n, 10))
	best := n
	bestI, bestJ := int64(0), int64(0)

	for i := 0; i < len(s); i++ {
		// build slice without the digit at index i
		rem := make([]byte, 0, len(s)-1)
		rem = append(rem, s[:i]...)
		rem = append(rem, s[i+1:]...)

		for j := 0; j <= len(rem); j++ {
			cand := make([]byte, 0, len(s))
			cand = append(cand, rem[:j]...)
			cand = append(cand, s[i])
			cand = append(cand, rem[j:]...)

			val, _ := strconv.ParseInt(string(cand), 10, 64)
			if val < best {
				best = val
				bestI = int64(i)
				bestJ = int64(j)
			}
		}
	}

	return [3]int64{best, bestI, bestJ}
}