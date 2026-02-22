/*
 * Generate All Chess960 Positions
 *
 * Background:
 * Freestyle chess, also known as Chess960, is a variant of chess where the
 * starting positions of the pieces are randomized.
 *
 * The pieces on the first rank can be placed anywhere, as long as the two
 * bishops are on different colored squares and the king is placed between
 * the two rooks. The 8th rank mirrors the 1st rank.
 *
 * There are 960 possible starting positions. Sorted lexicographically,
 * each integer from 1 to 960 maps to a specific position.
 *
 * Given a positive integer i (1-960), return a string representing the
 * board in the i'th Chess960 starting position.
 *
 * Pieces: Q, R, K, N, B, P (uppercase=white, lowercase=black)
 * "." for empty squares. "\n" terminates each row including the last.
 */
function chess960(i) {
}