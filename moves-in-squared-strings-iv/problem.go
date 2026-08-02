package main

// You are given a string of n lines, each substring being n characters long.
// For example: s = "abcd\nefgh\nijkl\nmnop"
//
// We will study some transformations of this square of strings:
//
// - Symmetry with respect to the antidiagonal: Diag2Sym
//   Diag2Sym("abcd\nefgh\nijkl\nmnop") => "plhd\nokgc\nnjfb\nmiea"
//
// - Counterclockwise rotation 90 degrees: Rot90Counter
//   Rot90Counter("abcd\nefgh\nijkl\nmnop") => "dhlp\ncgko\nbfjn\naeim"
//
// - SelfieDiag2Counterclock
//   It is initial string + string obtained by symmetry with respect to the antidiagonal + counterclockwise rotation 90 degrees.
//   SelfieDiag2Counterclock("abcd\nefgh\nijkl\nmnop") => "abcd|plhd|dhlp\nefgh|okgc|cgko\nijkl|njfb|bfjn\nmnop|miea|aeim"
//
// Task:
// Write the functions Diag2Sym, Rot90Counter, SelfieDiag2Counterclock
// and the high-order function Oper(fct, s) where fct is a function to apply to string s

func Diag2Sym(s string) string {
	return ""
}

func Rot90Counter(s string) string {
	return ""
}

func SelfieDiag2Counterclock(s string) string {
	return ""
}

func Oper(fct func(string) string, s string) string {
	return ""
}