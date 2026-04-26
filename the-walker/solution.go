package kata

import "math"

func Solve(a, b, c, alpha, beta, gamma int) []int {
	toRad := math.Pi / 180.0
	d1 := float64(alpha) * toRad
	d2 := (90.0 + float64(beta)) * toRad
	d3 := (180.0 + float64(gamma)) * toRad

	cx := float64(a)*math.Cos(d1) + float64(b)*math.Cos(d2) + float64(c)*math.Cos(d3)
	cy := float64(a)*math.Sin(d1) + float64(b)*math.Sin(d2) + float64(c)*math.Sin(d3)

	dist := math.Sqrt(cx*cx + cy*cy)
	angle := math.Atan2(cy, cx) * 180.0 / math.Pi
	if angle < 0 {
		angle += 360
	}

	distRound := int(math.Round(dist))
	deg := int(math.Floor(angle))
	minF := (angle - float64(deg)) * 60.0
	minI := int(math.Floor(minF))
	secI := int(math.Floor((minF - float64(minI)) * 60.0))

	return []int{distRound, deg, minI, secI}
}