/*
 * self_converge
 *
 * Goal: Given a number (with a minimum of 3 digits), return the number of
 * iterations it takes to arrive at a derived number that converges on to
 * itself, as per the Kaprekar routine.
 *
 * 0. Initialize a counter to count the number of iterations
 * 1. Take any number n, using at least two different digits.
 * 2. Arrange the digits in descending and ascending order to get two numbers,
 *    adding leading zeros if necessary so the width of the original is kept.
 * 3. Subtract the smaller from the bigger. Call this nseq.
 * 4. If nseq equals a previous value of n (cycle detected), return the count
 *    of iterations. Otherwise increment the counter and repeat with nseq.
 *    - If the sequence collapses to zero, return -1.
 *
 * Examples:
 *   1234  -> 4
 *   414   -> 5
 *   50000 -> 4
 */

function selfConverge(n) {
  // your code here
}