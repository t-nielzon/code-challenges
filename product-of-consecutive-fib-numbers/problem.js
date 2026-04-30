/*
 * The Fibonacci numbers are the numbers in the following integer sequence (Fn):
 * 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, ...
 *
 * such that:
 *   F(0) = 0
 *   F(1) = 1
 *   F(n) = F(n-1) + F(n-2)
 *
 * Given a number, say prod (for product), we search two Fibonacci numbers F(n) and F(n+1) verifying:
 *   F(n) * F(n+1) = prod
 *
 * Your function takes an integer (prod) and returns an array:
 *   - if F(n) * F(n+1) = prod: [F(n), F(n+1), true]
 *   - otherwise: [F(n), F(n+1), false] where F(n) is the smallest one such as F(n) * F(n+1) > prod
 *
 * Examples:
 *   productFib(714) --> [21, 34, true]
 *   productFib(800) --> [34, 55, false]
 */

function productFib(prod) {

}