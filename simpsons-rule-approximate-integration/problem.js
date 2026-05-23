/*
 * Simpson's Rule - Approximate Integration
 *
 * An integral ∫[a,b] f(x)dx can be approximated by Simpson's rule:
 *
 *   (b - a) / (3n) * ( f(a) + f(b)
 *                    + 4 * Σ_{i=1}^{n/2}   f(a + (2i - 1)h)
 *                    + 2 * Σ_{i=1}^{n/2-1} f(a + 2i*h) )
 *
 * where h = (b - a) / n and n is an even integer with a <= b.
 *
 * Apply Simpson's rule to f(x) = (3/2) * sin(x)^3 on the interval [0, pi].
 *
 * Write a function `simpson(n)` returning the approximation of the integral.
 *
 * Notes:
 *   - Do not round or truncate the result.
 *   - n is always even.
 *   - The exact value of the integral on [0, pi] is 2.
 */

function simpson(n) {
  // your code here
}