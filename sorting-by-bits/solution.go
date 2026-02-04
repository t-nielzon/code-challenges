package kata

import (
	"math/bits"
	"sort"
)

func SortByBit(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		bitsI := bits.OnesCount32(uint32(arr[i]))
		bitsJ := bits.OnesCount32(uint32(arr[j]))
		if bitsI != bitsJ {
			return bitsI < bitsJ
		}
		return arr[i] < arr[j]
	})
}