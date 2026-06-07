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

func temps(v0, slope, dTot float64) int {
	t, v, d, watts := 0.0, v0, 0.0, WATTS0

	// function(slope): slope is a grade percentage (rise/run), so the gravity
	// component along the incline is sin(atan(grade)).
	gravityComponent := GRAVITY_ACC * math.Sin(math.Atan(slope/100.0))

	for d < dTot {
		gamma := -gravityComponent - DRAG*math.Abs(v)*math.Abs(v)/MASS
		if watts > 0 && v > 0 {
			gamma += G_THRUST * watts / (v * MASS)
		}
		if math.Abs(gamma) <= 1e-5 {
			gamma = 0
		}

		v = v + gamma*DELTA_T
		if v-3.0 <= 1e-2 {
			return -1
		}

		d = d + v*DELTA_T/60.0
		watts = watts - D_WATTS*DELTA_T
		t = t + DELTA_T
	}

	return int(math.Round(t))
}