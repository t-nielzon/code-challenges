/*
 * Moves in squared strings (III)
 *
 * You are given a string of `n` lines, each substring being `n`
 * characters long. For example:
 *
 *   s = "abcd\nefgh\nijkl\nmnop"
 *
 * We will study some transformations of this square of strings.
 *
 * - Symmetry with respect to the main diagonal: diag_1_sym
 *     diag_1_sym(s) => "aeim\nbfjn\ncgko\ndhlp"
 *
 * - Clockwise rotation 90 degrees: rot_90_clock
 *     rot_90_clock(s) => "miea\nnjfb\nokgc\nplhd"
 *
 * - selfie_and_diag1(s)
 *   It is initial string + string obtained by symmetry with respect
 *   to the main diagonal.
 *     s = "abcd\nefgh\nijkl\nmnop" -->
 *     "abcd|aeim\nefgh|bfjn\nijkl|cgko\nmnop|dhlp"
 *
 * Task:
 * - Write the functions `diag_1_sym`, `rot_90_clock`, `selfie_and_diag1`
 * - Write the high-order function `oper(fct, s)` where `fct` is the
 *   function of one variable to apply to the string `s`.
 *
 * Examples:
 *   s = "abcd\nefgh\nijkl\nmnop"
 *   oper(diag_1_sym, s)      => "aeim\nbfjn\ncgko\ndhlp"
 *   oper(rot_90_clock, s)    => "miea\nnjfb\nokgc\nplhd"
 *   oper(selfie_and_diag1, s) => "abcd|aeim\nefgh|bfjn\nijkl|cgko\nmnop|dhlp"
 */

function diag_1_sym(s) {

}

function rot_90_clock(s) {

}

function selfie_and_diag1(s) {

}

function oper(fct, s) {

}