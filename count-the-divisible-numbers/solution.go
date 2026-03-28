package kata

func DivisibleCount(x, y, k uint64) uint64 {
	return y/k - (x-1)/k
}