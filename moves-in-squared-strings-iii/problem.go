package main

import "strings"

/*
Moves in squared strings (III)

You are given a string of n lines, each substring being n characters long.
We will study some transformations of this square of strings:

- Symmetry with respect to the main diagonal: diag_1_sym
  diag_1_sym(s) => "aeim\nbfjn\ncgko\ndhlp"

- Clockwise rotation 90 degrees: rot_90_clock
  rot_90_clock(s) => "miea\nnjfb\nokgc\nplhd"

- selfie_and_diag1(s)
  It is initial string + string obtained by symmetry with respect to the main diagonal.
  selfie_and_diag1(s) => "abcd|aeim\nefgh|bfjn\nijkl|cgko\nmnop|dhlp"

Task:
Write these functions and a high-order function oper(fct, s) where
fct is the function to apply to the string s.
*/

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