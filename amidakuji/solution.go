package kata

func Amidakuji(ladder []string) []int {
	if len(ladder) == 0 {
		return []int{}
	}
	w := len(ladder[0])
	res := make([]int, w)
	for i := range res {
		res[i] = i
	}
	for _, row := range ladder {
		for i := 0; i < w-1; i++ {
			if row[i] == '1' {
				res[i], res[i+1] = res[i+1], res[i]
			}
		}
	}
	return res
}