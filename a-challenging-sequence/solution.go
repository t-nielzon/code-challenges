package kata

import "sort"

func HashRadSeq(nMax, k int) int {
	rad := make([]int, nMax+1)
	for i := 0; i <= nMax; i++ {
		rad[i] = 1
	}
	for i := 2; i <= nMax; i++ {
		if rad[i] == 1 {
			for j := i; j <= nMax; j += i {
				rad[j] *= i
			}
		}
	}

	type pair struct{ n, r int }
	arr := make([]pair, nMax)
	for i := 1; i <= nMax; i++ {
		arr[i-1] = pair{i, rad[i]}
	}
	sort.Slice(arr, func(a, b int) bool {
		if arr[a].r != arr[b].r {
			return arr[a].r < arr[b].r
		}
		return arr[a].n < arr[b].n
	})
	return arr[k-1].n
}