/*
 * Easy Cyclist's Training
 *
 * John has bought a bike and wants to simulate his mountain biking trip.
 *
 * His trip consists of an ascent of dTot km with an average slope (%).
 * No wind, constant mass (80 kg), initial power 225 watts.
 * No rolling resistance. Pushed start at v0 km/h. Initial acceleration is 0.
 * Time step DELTA_T = 1/60 min.
 *
 * He loses D_WATTS * DELTA_T of power each step.
 *
 * Acceleration has three components:
 *   1) gravity:  -GRAVITY_ACC * sin(atan(slope/100))
 *   2) air drag: -DRAG * |v|^2 / MASS
 *   3) thrust (if watts > 0 and v > 0): +G_THRUST * watts / (v * MASS)
 *   4) if |acceleration| <= 1e-5, set to 0
 *
 * If v - 3.0 <= 1e-2, John gives up (return -1).
 *
 * speed(t+dt) = speed(t) + gamma * DELTA_T
 * dist(t+dt)  = dist(t) + speed(t+dt) * DELTA_T / 60
 *
 * Return rounded integer time in minutes, or -1 if he gives up.
 *
 * Examples:
 *   temps(30, 5, 30)  -> 114
 *   temps(30, 20, 30) -> -1
 *   temps(30, 8, 20)  -> 110
 */
function temps(v0, slope, dTot) {
}