package kata

func FindInArray(seq []interface{}, fn func(interface{}, int) bool) int {
	for i, v := range seq {
		if fn(v, i) {
			return i
		}
	}
	return -1
}