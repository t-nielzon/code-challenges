package kata

func PaperFold() func() int {
	n := 0
	return func() int {
		n++
		m := n
		for m&1 == 0 {
			m >>= 1
		}
		if m&2 == 0 {
			return 1
		}
		return 0
	}
}