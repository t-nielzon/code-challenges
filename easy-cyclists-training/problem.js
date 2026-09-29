/*
 * Easy Cyclist's Training (Difficulty: 5 kyu)
 * 
 * John has bought a bike and wants to simulate climbing a mountain.
 * 
 * Given Information:
 * - Trip: ascent of dTot kilometers with average slope of slope percent
 * - John's mass: 80 kg (constant)
 * - Initial power: 225 watts
 * - No wind, no rolling resistance
 * - Initial speed: v0 km/h
 * - Initial acceleration: 0
 * - Time step: DELTA_T = 1/60 minutes
 * 
 * Physics:
 * - Speed update: v(t+DELTA_T) = v(t) + gamma * DELTA_T
 * - Distance update: d(t+DELTA_T) = d(t) + v(t+DELTA_T) * DELTA_T / 60.0
 * - Power loss: watts(t+DELTA_T) = watts(t) - D_WATTS * DELTA_T
 * 
 * Acceleration (gamma in km/h/min) has three components:
 * 1. Gravity: -GRAVITY_ACC * (slope / 100)
 * 2. Air drag: -DRAG * abs(v)^2 / MASS
 * 3. Thrust (if watts > 0 and v > 0): +G_THRUST * watts / (v * MASS)
 * 
 * If abs(gamma) < 1e-5, set gamma to 0.
 * If v - 3.0 <= 1e-2, John gives up and return -1.
 * 
 * Constants:
 * - GRAVITY_ACC = 9.81 * 3.6 * 60.0
 * - DRAG = 60.0 * 0.3 / 3.6
 * - DELTA_T = 1.0 / 60.0
 * - G_THRUST = 60 * 3.6 * 3.6
 * - MASS = 80.0
 * - WATTS0 = 225.0
 * - D_WATTS = 0.5
 * 
 * @param {number} v0 - Initial speed in km/h
 * @param {number} slope - Ascent in percentage
 * @param {number} dTot - Distance to travel in km
 * @return {number} Time in minutes (rounded integer), or -1 if John gives up
 * 
 * Examples:
 * temps(30, 5, 30) -> 114
 * temps(30, 20, 30) -> -1
 * temps(30, 8, 20) -> 110
 */

function temps(v0, slope, dTot) {
  
}