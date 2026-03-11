package kata

func PartsSums(ls []uint64) []uint64 {
	result := make([]uint64, len(ls)+1)
	var total uint64
	for _, v := range ls {
		total += v
	}
	result[0] = total
	for i, v := range ls {
		total -= v
		result[i+1] = total
	}
	return result
}