/*
 * Simpson's Rule - Approximate Integration
 *
 * An integral ∫[a,b] f(x)dx can be approximated by Simpson's rule:
 *
 *   (b - a) / (3n) * ( f(a) + f(b)
 *                    + 4 * Σ[i=1..n/2]   f(a + (2i - 1)h)
 *                    + 2 * Σ[i=1..n/2-1] f(a + 2i*h) )
 *
 * where h = (b - a) / n, n is an even integer, and a <= b.
 *
 * Apply Simpson's rule to:
 *   f(x) = (3/2) * sin(x)^3
 *
 * Write a function `simpson(n)` that returns the value of the integral
 * of f on the interval [0, pi].
 *
 * Notes:
 *  - Do not round or truncate results.
 *  - n is always even.
 *  - The exact value of the integral on [0, pi] is 2.
 */

function simpson(n) {

}