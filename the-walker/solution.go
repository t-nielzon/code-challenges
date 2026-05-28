package kata

import "math"

func Solve(a, b, c, alpha, beta, gamma int) []int {
	const pi = 3.14159265358979323846

	aR := float64(alpha) * pi / 180
	bR := float64(beta) * pi / 180
	gR := float64(gamma) * pi / 180

	cx := float64(a)*math.Cos(aR) - float64(b)*math.Sin(bR) - float64(c)*math.Cos(gR)
	cy := float64(a)*math.Sin(aR) + float64(b)*math.Cos(bR) - float64(c)*math.Sin(gR)

	dist := math.Sqrt(cx*cx + cy*cy)
	angle := math.Atan2(cy, cx) * 180 / pi
	if angle < 0 {
		angle += 360
	}

	distInt := int(math.Round(dist))
	deg := int(math.Floor(angle))
	minF := (angle - float64(deg)) * 60
	minInt := int(math.Floor(minF))
	sec := int(math.Floor((minF - float64(minInt)) * 60))

	return []int{distInt, deg, minInt, sec}
}