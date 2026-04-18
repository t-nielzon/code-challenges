/*
Given

- a semi-inclusive interval `I = [l, u)` (l is in interval I but u is not)
  `l` and `u` being floating numbers `(0 <= l < u)`,

- an integer `n (n > 0)`
- a function `f: x (float number) -> f(x) (float number)`

we want to return as a list the `n` values:

`f(l), f(l + 1 * d), ..., f(u -d)` where `d = (u - l) / n`

Call this function `interp`:

`interp(f, l, u, n) -> [f(l), f(l + 1 * d), ..., f(u - d)]`

The `n` resulting values will be floored to two decimals.
Use: floor(y * 100.0) / 100.0
*/

function interp(f, l, u, n) {

}