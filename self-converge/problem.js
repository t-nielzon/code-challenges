/*
 * self_converge
 *
 * Goal: Given a number (with a minimum of 3 digits), return the number of
 * iterations it takes to arrive at a derived number that converges on to
 * itself, as per the following Kaprekar routine. As a learning exercise,
 * come up with a solution that uses recursion.
 *
 * Steps:
 * 0. Initialize a counter to count the number of iterations
 * 1. Take any number n, using at least two different digits.
 * 2. Arrange the digits in descending and then in ascending order to get
 *    two numbers, adding leading zeros if necessary to maintain original width.
 * 3. Subtract the smaller number from the bigger number. Let us call this nseq.
 * 4. Check if nseq equals any previous value of n. If not, increment the
 *    iteration counter and go back to step 2.
 *
 * If the sequence collapses to zero, return -1.
 *
 * Example:
 *   1234  -> 4
 *   414   -> 5
 *   50000 -> 4
 */
function selfConverge(n) {
}