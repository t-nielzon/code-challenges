package kata

// two-pointer sweep keeps it linear: water above any column is bounded by the
// smaller of the tallest walls seen from each side, and we always advance the
// side whose running max is provably the limiting one.
func Flood(towers []int) int {
	left, right := 0, len(towers)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if towers[left] < towers[right] {
			if towers[left] >= leftMax {
				leftMax = towers[left]
			} else {
				water += leftMax - towers[left]
			}
			left++
		} else {
			if towers[right] >= rightMax {
				rightMax = towers[right]
			} else {
				water += rightMax - towers[right]
			}
			right--
		}
	}

	return water
}