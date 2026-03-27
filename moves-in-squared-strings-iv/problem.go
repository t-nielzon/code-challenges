/*
You are given a string of n lines, each substring being n characters long.

Transformations:
- diag_2_sym: Symmetry with respect to the main cross diagonal
- rot_90_counter: Counterclockwise rotation 90 degrees
- selfie_diag2_counterclock: initial + diag_2_sym + rot_90_counter joined by |
- oper(fct, s): high-order function applying fct to s

Example:
  s = "abcd\nefgh\nijkl\nmnop"
  oper(diag_2_sym, s) => "plhd\nokgc\nnjfb\nmiea"
  oper(rot_90_counter, s) => "dhlp\ncgko\nbfjn\naeim"
  oper(selfie_diag2_counterclock, s) => "abcd|plhd|dhlp\nefgh|okgc|cgko\nijkl|njfb|bfjn\nmnop|miea|aeim"
*/
package kata

func Diag2Sym(s string) string {
}

func Rot90Counter(s string) string {
}

func SelfieDiag2Counterclock(s string) string {
}

func Oper(fct func(string) string, s string) string {
}