/*
 * Physics - Mechanics - Atwood Machines in Series
 *
 * The Atwood machine consists of a weightless, frictionless pulley with two
 * masses suspended from a weightless string at each end. One of these masses
 * can be replaced by another Atwood machine, and again, and again...
 *
 * Write a function that receives a list of 2 to 100 masses (kg) and returns
 * the acceleration (m/s/s) of the first mass.
 *
 * Physics Reference:
 * - Force = Mass x Acceleration
 * - Gravitational Force (N) = Mass (kg) x 9.80665 (m/s/s)
 * - Tension Force is equal for objects on two ends of the same string
 * - Find acceleration by combining all forces on an object
 * - In a many-mass system, multiple masses can often be reduced to one
 *   "effective" mass
 * - Solve the two-mass system first, then replace the second mass with a
 *   pulley (recurse)
 *
 * Key equations:
 *   a1 = g * (m2 - m1) / (m1 + m2)
 *   Effective Lower Pulley Mass = 4 * m2 * m3 / (m2 + m3)
 */

function atwood(masses) {

}