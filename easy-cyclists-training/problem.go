/*
John has bought a bike but before going moutain biking he wants us to do a few simulations.

He gathered information:
- His trip will consist of an ascent of dTot kilometers with an average slope of slope percent
- John's mass is constant MASS = 80 (kg), initial power WATTS0 = 225 (watts)
- No wind, no rolling resistance
- Initial speed v0 (km/h), initial acceleration gamma = 0
- DELTA_T = 1.0 / 60.0 (in minutes)
- He loses D_WATTS * DELTA_T of power at each DELTA_T
- Acceleration components:
    1) -GRAVITY_ACC * function(slope) where slope is a percentage
    2) -DRAG * abs(v) * abs(v) / MASS
    3) if watts and v are strictly positive: +G_THRUST * watts / (v * MASS)
    4) if total acceleration is <= 1e-5 set it to 0
- If v - 3.0 <= 1e-2 John gives up

Constants:
GRAVITY_ACC = 9.81 * 3.6 * 60.0
DRAG        = 60.0 * 0.3 / 3.6
DELTA_T     = 1.0 / 60.0
G_THRUST    = 60 * 3.6 * 3.6
MASS        = 80.0
WATTS0      = 225.0
D_WATTS     = 0.5

Task:
Write function temps(v0, slope, dTot) which returns as a rounded integer
the time t in minutes needed to travel dTot. If John gives up return -1.

Examples:
temps(30, 5, 30)  -> 114
temps(30, 20, 30) -> -1
temps(30, 8, 20)  -> 110
*/

package kata

func Temps(v0, slope, dTot float64) int {
}