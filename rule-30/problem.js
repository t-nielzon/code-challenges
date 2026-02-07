/*
 * Rule 30
 *
 * Rule 30 is a one-dimensional binary cellular automaton.
 *
 * Complete the function that takes as input an array of 0s and 1s and a
 * non-negative integer n that represents the number of iterations. This
 * function has to perform the nth iteration of Rule 30 with the given input.
 *
 * The rule to derive a cell from itself and its neighbour is:
 *
 * Current cell  | 000 | 001 | 010 | 011 | 100 | 101 | 110 | 111
 * New cell      |  0  |  1  |  1  |  1  |  1  |  0  |  0  |  0
 *
 * The new state of a certain cell depends on its neighborhood. In Current cell
 * you have the nth cell with its left and right neighbor.
 *
 * Borders of the list are always 0.
 * Return an array of 0 and 1.
 *
 * Example: starting from [1] and iterating 5 times:
 * - Add borders: [0, 1, 0]
 * - Apply rule 30: [1, 1, 1] (first iteration)
 * - After 5 iterations: [1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1]
 */
function rule30(cells, n) {
}

module.exports = { rule30 };