/*
 * # Task
 * Given a string representing a simple fraction `x/y`, your function must
 * return a string representing the corresponding mixed fraction in the
 * following format:
 *
 *   [sign]a b/c
 *
 * where `a` is integer part and `b/c` is irreducible proper fraction. There
 * must be exactly one space between `a` and `b/c`. Provide `[sign]` only if
 * negative (and non zero) and only at the beginning of the number (both
 * integer part and fractional part must be provided absolute).
 *
 * If the `x/y` equals the integer part, return integer part only. If integer
 * part is zero, return the irreducible proper fraction only. In both of these
 * cases, the resulting string must not contain any spaces.
 *
 * Division by zero should raise an error.
 *
 * # Value ranges
 *   -10 000 000 < x < 10 000 000
 *   -10 000 000 < y < 10 000 000
 *
 * # Examples
 *   `42/9`      -> `4 2/3`
 *   `6/3`       -> `2`
 *   `4/6`       -> `2/3`
 *   `0/18891`   -> `0`
 *   `-10/7`     -> `-1 3/7`
 *   `0/0`, `3/0` -> raise a zero division error
 *
 * # Note
 * Make sure not to modify the input of your function in-place.
 */

function convertToMixedNumber(input) {
  // your code here
}