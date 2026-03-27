/*
 * Coding with Squared Strings
 *
 * A squared string has n lines, each substring being n characters long.
 * e.g. s = "abcd\nefgh\nijkl\nmnop" is a squared string of size 4.
 *
 * Coding:
 * - Take text t of length l
 * - Find smallest n where n*n >= l
 * - Pad t with char code 11 until length is n*n
 * - Form squared string of size n
 * - Apply clockwise 90° rotation
 *
 * Decoding:
 * - Take squared string s from coding
 * - Apply counter-clockwise 90° rotation
 * - Clean padding to recover original text
 */

function code(s) {
}

function decode(s) {
}