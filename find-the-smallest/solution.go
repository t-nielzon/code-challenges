package kata

import "strconv"

func Smallest(n int) [3]int {
	s := strconv.Itoa(n)
	best := n
	bi, bj := 0, 0
	first := true
	for i := 0; i < len(s); i++ {
		rem := s[:i] + s[i+1:]
		d := s[i]
		for j := 0; j <= len(rem); j++ {
			cand := rem[:j] + string(d) + rem[j:]
			v, err := strconv.Atoi(cand)
			if err != nil {
				continue
			}
			if first || v < best || (v == best && (i < bi || (i == bi && j < bj))) {
				best = v
				bi = i
				bj = j
				first = false
			}
		}
	}
	return [3]int{best, bi, bj}
}