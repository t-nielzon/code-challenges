package main

import "math"

func AreaOfPolygonInsideCircle(circleRadius float64, numberOfSides int) float64 {
	// area of regular polygon = (n * r^2 * sin(2*pi/n)) / 2
	area := float64(numberOfSides) * circleRadius * circleRadius * math.Sin(2*math.Pi/float64(numberOfSides)) / 2
	// round to 3 decimal places
	return math.Round(area*1000) / 1000
}