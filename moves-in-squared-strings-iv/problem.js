/*
 * Moves in squared strings (IV)
 *
 * You are given a string of `n` lines, each substring being `n` characters long.
 * For example: s = "abcd\nefgh\nijkl\nmnop"
 *
 * We study some transformations of this square of strings.
 *
 * - Symmetry with respect to the antidiagonal: diag_2_sym
 *     diag_2_sym(s) => "plhd\nokgc\nnjfb\nmiea"
 *
 * - Counterclockwise rotation 90 degrees: rot_90_counter
 *     rot_90_counter(s) => "dhlp\ncgko\nbfjn\naeim"
 *
 * - selfie_diag2_counterclock
 *     It is initial string + string obtained by symmetry with respect to the
 *     anti diagonal + counterclockwise rotation 90 degrees.
 *     s = "abcd\nefgh\nijkl\nmnop" -->
 *     "abcd|plhd|dhlp\nefgh|okgc|cgko\nijkl|njfb|bfjn\nmnop|miea|aeim"
 *
 * Task:
 * - Write functions diag_2_sym, rot_90_counter, selfie_diag2_counterclock
 * - Write high-order function oper(fct, s) where fct is the function of one
 *   variable to apply to the string s.
 */

function diag_2_sym(s) {

}

function rot_90_counter(s) {

}

function selfie_diag2_counterclock(s) {

}

function oper(fct, s) {

}