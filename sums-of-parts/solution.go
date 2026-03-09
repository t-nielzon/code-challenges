package kata

func PartsSums(ls []uint64) []uint64 {
	n := len(ls)
	result := make([]uint64, n+1)
	var total uint64
	for _, v := range ls {
		total += v
	}
	result[0] = total
	for i := 0; i < n; i++ {
		total -= ls[i]
		result[i+1] = total
	}
	return result
}