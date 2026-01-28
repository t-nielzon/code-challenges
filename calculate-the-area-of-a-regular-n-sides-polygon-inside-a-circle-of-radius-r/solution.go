package kata

import "math"

func AreaOfPolygonInsideCircle(circleRadius float64, numberOfSides int) float64 {
	n := float64(numberOfSides)
	area := 0.5 * n * circleRadius * circleRadius * math.Sin(2*math.Pi/n)
	return math.Round(area*1000) / 1000
}