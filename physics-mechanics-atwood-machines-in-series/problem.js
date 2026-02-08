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
 * Key equations:
 *   a1 = g * (m2 - m1) / (m1 + m2)
 *   Effective mass of a sub-pulley = 4 * m_a * m_b / (m_a + m_b)
 */

function atwoodMachines(masses) {
}

module.exports = atwoodMachines;