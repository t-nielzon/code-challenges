/*
Consider the sequence S(n, z) = (1 - z)(z + z^2 + z^3 + ... + z^n) where z is a complex number
and n a positive integer (n > 0).

When n goes to infinity and z is in the domain of convergence D, S(n, z) goes to a finite limit lim.

Determine the smallest integer n such that abs(S(n, z) - lim) < eps.
If z is outside D, return -1.

z is represented by two reals x (real part) and y (imaginary part).

Examples:
  f(0.3, 0.5, 1e-4) -> 17
  f(30, 5, 1e-4)    -> -1
*/

function f(x, y, eps) {
  // your code
}