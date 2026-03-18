/*
 * Chess - underpromote?
 *
 * Max wants to improve his chess skills. To achieve his goal he decided to do exercises.
 * Help Max at his (K)nightmares by solving this Kata.
 *
 * Task:
 * Return the minimum N (amount of moves) needed on an 8x8 board starting with a white pawn
 * from initial position start ("a2" to "h7", 1st and 8th rank excluded) to target position
 * target ("a1" to "h8"). If start and target are equal then return 0.
 *
 * Notes:
 * - No other pieces on the board except the pawn on start
 * - A pawn can only move one square up
 * - A pawn on the 2nd rank cannot move 2 squares
 * - If the pawn reaches the 8th rank it will be promoted to a Queen or a Knight.
 *   It is not required to promote if the pawn can reach the target only by moving forward.
 * - The Queen can move in all directions with no limits but cannot imitate a Knight's movement.
 */

function underpromote(start, target) {
}