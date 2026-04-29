package kata

func CitySwim(towers []int) int {
	n := len(towers)
	if n < 3 {
		return 0
	}
	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	total := 0
	for left < right {
		if towers[left] < towers[right] {
			if towers[left] >= leftMax {
				leftMax = towers[left]
			} else {
				total += leftMax - towers[left]
			}
			left++
		} else {
			if towers[right] >= rightMax {
				rightMax = towers[right]
			} else {
				total += rightMax - towers[right]
			}
			right--
		}
	}
	return total
}