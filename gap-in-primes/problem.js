/*
 * Gap in Primes
 *
 * The prime numbers are not regularly spaced. For example from 2 to 3 the
 * gap is 1. From 3 to 5 the gap is 2. From 7 to 11 it is 4.
 *
 * A prime gap of length n is a run of n-1 consecutive composite numbers
 * between two successive primes.
 *
 * Write a function gap with parameters:
 *  - g (integer >= 2): the gap we are looking for
 *  - m (integer > 2): the start of the search (m inclusive)
 *  - n (integer >= m): the end of the search (n inclusive)
 *
 * The function should return the first pair of two successive prime numbers
 * spaced with a gap of g between the limits m and n if they exist, otherwise
 * null.
 *
 * Examples:
 *  gap(2, 3, 50)     --> [3, 5]
 *  gap(2, 5, 7)      --> [5, 7]
 *  gap(2, 5, 5)      --> null
 *  gap(4, 130, 200)  --> [163, 167]
 *  gap(6, 100, 110)  --> null
 */

function gap(g, m, n) {
  // your code here
}