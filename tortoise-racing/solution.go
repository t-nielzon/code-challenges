package kata

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}
	totalSeconds := g * 3600 / (v2 - v1)
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60
	return [3]int{hours, minutes, seconds}
}