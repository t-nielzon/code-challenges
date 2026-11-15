package main

func ballUpwards(v0 float64) int {
	const g = 9.81
	v0_ms := v0 / 3.6 // convert km/h to m/s

	maxHeight := 0.0
	maxTime := 0

	for t := 0; ; t++ {
		actualTime := float64(t) / 10.0 // convert tenths to seconds
		height := v0_ms*actualTime - 0.5*g*actualTime*actualTime

		if height < 0 {
			break
		}

		if height > maxHeight {
			maxHeight = height
			maxTime = t
		}
	}

	return maxTime
}