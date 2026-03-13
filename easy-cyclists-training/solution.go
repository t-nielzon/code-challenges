package kata

import "math"

const (
	GRAVITY_ACC = 9.81 * 3.6 * 60.0
	DRAG        = 60.0 * 0.3 / 3.6
	DELTA_T     = 1.0 / 60.0
	G_THRUST    = 60 * 3.6 * 3.6
	MASS        = 80.0
	WATTS0      = 225.0
	D_WATTS     = 0.5
)

func Temps(v0, slope, dTot float64) int {
	t := 0.0
	v := v0
	d := 0.0
	watts := WATTS0
	slopeRad := math.Sin(math.Atan(slope / 100.0))

	for d < dTot {
		gamma := -GRAVITY_ACC * slopeRad
		gamma -= DRAG * math.Abs(v) * math.Abs(v) / MASS
		if watts > 0 && v > 0 {
			gamma += G_THRUST * watts / (v * MASS)
		}
		if gamma <= 1e-5 {
			gamma = 0
		}

		v = v + gamma*DELTA_T
		d = d + v*DELTA_T/60.0
		t += DELTA_T
		watts -= D_WATTS * DELTA_T

		if v-3.0 <= 1e-2 {
			return -1
		}
	}

	return int(math.Round(t))
}