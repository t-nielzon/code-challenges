package kata

import "math/bits"

func NextHigher(n int) int {
	target := bits.OnesCount(uint(n))
	for i := n + 1; ; i++ {
		if bits.OnesCount(uint(i)) == target {
			return i
		}
	}
}