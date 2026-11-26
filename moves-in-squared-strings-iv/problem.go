package kata

/*
Moves in squared strings (IV)

You are given a string of n lines, each substring being n characters long.
For example: s = "abcd\nefgh\nijkl\nmnop"

We will study some transformations of this square of strings.

- Symmetry with respect to the antidiagonal: Diag2Sym
  Diag2Sym(s) => "plhd\nokgc\nnjfb\nmiea"

- Counterclockwise rotation 90 degrees: Rot90Counter
  Rot90Counter(s) => "dhlp\ncgko\nbfjn\naeim"

- SelfieDiag2Counterclock
  It is initial string + string obtained by symmetry with respect to the antidiagonal + counterclockwise rotation 90 degrees.
  s = "abcd\nefgh\nijkl\nmnop" -->
  "abcd|plhd|dhlp\nefgh|okgc|cgko\nijkl|njfb|bfjn\nmnop|miea|aeim"

Task:
- Write functions Diag2Sym, Rot90Counter, SelfieDiag2Counterclock
- Write higher-order function Oper(fct, s) where fct is a function to apply to string s
*/

func Diag2Sym(s string) string {
	// stub
	return ""
}

func Rot90Counter(s string) string {
	// stub
	return ""
}

func SelfieDiag2Counterclock(s string) string {
	// stub
	return ""
}

type TransformFunc func(string) string

func Oper(fct TransformFunc, s string) string {
	// stub
	return ""
}