package main

func race(v1, v2, g int) *[3]int {
	if v1 >= v2 {
		return nil
	}

	// calculate total seconds for B to catch up: gap / closing speed * seconds per hour
	seconds := (g * 3600) / (v2 - v1)

	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60

	return &[3]int{hours, minutes, secs}
}