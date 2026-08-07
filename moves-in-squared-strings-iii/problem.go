/*
Moves in squared strings (III)

You are given a string of n lines, each substring being n characters long.
For example: s = "abcd\nefgh\nijkl\nmnop"

We will study some transformations of this square of strings:

1. Symmetry with respect to the main diagonal: diag_1_sym
   diag_1_sym(s) => "aeim\nbfjn\ncgko\ndhlp"

2. Clockwise rotation 90 degrees: rot_90_clock
   rot_90_clock(s) => "miea\nnjfb\nokgc\nplhd"

3. selfie_and_diag1
   It is initial string + string obtained by symmetry with respect to the main diagonal.
   selfie_and_diag1(s) => "abcd|aeim\nefgh|bfjn\nijkl|cgko\nmnop|dhlp"

Task:
- Write functions diag_1_sym, rot_90_clock, selfie_and_diag1
- Write high-order function oper(fct, s) where fct is one of the above functions

Examples:
- oper(diag_1_sym, s) => "aeim\nbfjn\ncgko\ndhlp"
- oper(rot_90_clock, s) => "miea\nnjfb\nokgc\nplhd"
- oper(selfie_and_diag1, s) => "abcd|aeim\nefgh|bfjn\nijkl|cgko\nmnop|dhlp"
*/
package kata

func diag_1_sym(s string) string {
	// TODO: implement
	return ""
}

func rot_90_clock(s string) string {
	// TODO: implement
	return ""
}

func selfie_and_diag1(s string) string {
	// TODO: implement
	return ""
}

func oper(fct func(string) string, s string) string {
	// TODO: implement
	return ""
}