/*
 * Moves in squared strings (III)
 * 
 * You are given a string of n lines, each substring being n characters long.
 * For example: s = "abcd\nefgh\nijkl\nmnop"
 * 
 * Transformations:
 * - diag_1_sym: Symmetry with respect to the main diagonal (transpose)
 *   diag_1_sym(s) => "aeim\nbfjn\ncgko\ndhlp"
 * 
 * - rot_90_clock: Clockwise rotation 90 degrees
 *   rot_90_clock(s) => "miea\nnjfb\nokgc\nplhd"
 * 
 * - selfie_and_diag1: Initial string + diagonal symmetry, separated by |
 *   selfie_and_diag1(s) => "abcd|aeim\nefgh|bfjn\nijkl|cgko\nmnop|dhlp"
 * 
 * Task: Write these three functions and a higher-order function oper(fct, s)
 * that applies the given function to the string s.
 */

function diag_1_sym(s) {
}

function rot_90_clock(s) {
}

function selfie_and_diag1(s) {
}

function oper(fct, s) {
}