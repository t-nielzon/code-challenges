package main

import "math"

func BallUpwards(v0 int) int {
	// convert velocity from km/h to m/s
	v0_ms := float64(v0) * 5 / 18
	g := 9.81

	// the maximum height occurs at t = v0 / g (in seconds)
	t_max_seconds := v0_ms / g

	// convert to tenths of seconds and round to nearest recording time
	return int(math.Round(t_max_seconds * 10))
}