/*
 * PI approximation
 *
 * The aim of the kata is to try to show how difficult it can be to calculate
 * decimals of an irrational number with a certain precision. We have chosen
 * to get a few decimals of the number "pi" using the following infinite series
 * (Leibniz 1646–1716):
 *
 * PI / 4 = 1 - 1/3 + 1/5 - 1/7 + ...
 *
 * Given as input a precision of epsilon, we stop the iterative process when
 * the absolute value of the difference between our calculation using the
 * Leibniz series and Math.PI is less than epsilon.
 *
 * Returns an array with:
 * - the number of iterations
 * - the approximation of PI with 10 decimals
 *
 * Example: iterPi(0.001) --> [1000, 3.1405926538]
 */
function iterPi(epsilon) {
}

module.exports = { iterPi };