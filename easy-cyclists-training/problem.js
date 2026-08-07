/*
John has bought a bike and wants to simulate his mountain climbing performance.

Trip parameters:
- Distance to climb: dTot (km)
- Slope: slope (percent)
- Initial speed: v0 (km/h)

Constants:
GRAVITY_ACC = 9.81 * 3.6 * 60.0
DRAG = 60.0 * 0.3 / 3.6
DELTA_T = 1.0 / 60.0
G_THRUST = 60 * 3.6 * 3.6
MASS = 80.0
WATTS0 = 225.0
D_WATTS = 0.5

Acceleration components:
1) Gravity: - GRAVITY_ACC * sin(arctan(slope / 100))
2) Air drag: - DRAG * v^2 / MASS
3) Thrust: + G_THRUST * watts / (v * MASS) (if v > 0 and watts > 0)
4) If |acceleration| <= 1e-5, set to 0

John gives up if v - 3.0 <= 1e-2

Task:
Write function temps(v0, slope, dTot) which returns the time in minutes to travel dTot km.
Return -1 if John gives up.

Examples:
temps(30, 5, 30) -> 114
temps(30, 20, 30) -> -1
temps(30, 8, 20) -> 110
*/

function temps(v0, slope, dTot) {
  
}