package kata

import "sort"

func Solve(arr []int) []int {
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

	result := make([]int, len(arr))
	copy(result, arr)

	sort.SliceStable(result, func(i, j int) bool {
		fi, fj := freq[result[i]], freq[result[j]]
		if fi != fj {
			return fi > fj
		}
		return result[i] < result[j]
	})

	return result
}