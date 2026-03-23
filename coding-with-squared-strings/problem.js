/*
 * Coding with Squared Strings
 *
 * A squared string has n lines, each substring being n characters long.
 * e.g. s = "abcd\nefgh\nijkl\nmnop" is a squared string of size 4.
 *
 * Coding:
 * - Given text t of length l, find smallest n where n*n >= l
 * - Pad t with char code 11 until length is n*n
 * - Form squared string of size n
 * - Apply clockwise 90° rotation
 *
 * Decoding:
 * - Apply counter-clockwise 90° rotation
 * - Remove padding chars (char code 11)
 *
 * Example:
 * t = "I.was.going.fishing.that.morning.at.ten.o'clock"
 * code(t) -> "c.nhsoI\nltiahi.\noentinw\ncng.nga\nk..mg.s\n\voao.f.\n\v'trtig"
 * decode(code(t)) == "I.was.going.fishing.that.morning.at.ten.o'clock"
 */

function code(s) {
}

function decode(s) {
}