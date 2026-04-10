/*
 * self_converge
 *
 * Goal: Given a number (with a minimum of 3 digits), return the number
 * of iterations it takes to arrive at a derived number that converges
 * on to itself, as per the following Kaprekar routine. Use recursion.
 *
 * 0. Initialize a counter to count the number of iterations
 * 1. Take any four-digit number n, using at least two different digits.
 * 2. Arrange the digits in descending and then in ascending order to get
 *    two four-digit numbers, adding leading zeros if necessary.
 *    - Add as many zeroes so that the width of the original number is maintained.
 * 3. Subtract the smaller number from the bigger number. Call this nseq.
 * 4. Check if nseq equals a previous value of n.
 *    - If not, increment the iteration counter and go back to step 2.
 *    - If the sequence collapses to zero, return -1.
 *
 * For 5+ digit numbers, convergence occurs on a cycle. Detect this cycle
 * by comparing to all previous values of n.
 *
 * Examples:
 *   1234  -> 4
 *   414   -> 5
 *   50000 -> 4
 */
function selfConverge(n) {
}