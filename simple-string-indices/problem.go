/*
In this Kata, you will be given a string with brackets and an index of an opening bracket
and your task will be to return the index of the matching closing bracket.

Both the input and returned index are 0-based except in Fortran where it is 1-based.
An opening brace will always have a closing brace. Return -1 if there is no answer
(in Haskell, return Nothing; in Fortran, return 0; in Go, return an error).

Examples:
  solve("((1)23(45))(aB)", 0)  = 10
  solve("((1)23(45))(aB)", 1)  = 3
  solve("((1)23(45))(aB)", 2)  = error (no opening bracket at index 2)
  solve("((1)23(45))(aB)", 6)  = 9
  solve("((1)23(45))(aB)", 11) = 14
  solve("((>)|?(*'))(yZ)", 11) = 14

Input consists of letters, numbers and special characters, but no spaces.
The only brackets will be ( and ).
*/

package kata

func Solve(s string, idx int) (int, error) {
}