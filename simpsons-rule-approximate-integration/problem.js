/*
 * Simpson's Rule - Approximate Integration
 *
 * An integral of f(x) from a to b can be approximated by Simpson's rule:
 *
 *   (b - a) / (3n) * (f(a) + f(b)
 *       + 4 * sum_{i=1..n/2} f(a + (2i - 1) * h)
 *       + 2 * sum_{i=1..n/2-1} f(a + 2i * h))
 *
 * where h = (b - a) / n and n is an even integer with a <= b.
 *
 * Apply Simpson's rule to:
 *
 *   f(x) = (3 / 2) * sin(x)^3
 *
 * on the interval [0, pi].
 *
 * Write a function `simpson(n)` returning the approximation.
 * The exact value of the integral is 2.
 *
 * Notes:
 *  - Don't round or truncate results.
 *  - n is always even.
 */

function simpson(n) {

}