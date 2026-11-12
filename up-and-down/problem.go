/*
Kata: up AND down
Difficulty: 6 kyu

You are given a string s made up of substrings separated by whitespaces.
Return a string t with lengths following the pattern:
length t(0) <= length t(1) >= length t(2) <= length t(3) >= length t(4) ...

Process from left to right, moving only adjacent (consecutive) strings when needed.
Minimize the number of moves.

Once the pattern is achieved, format so that:
- even indices (0, 2, 4, ...) are lowercase
- odd indices (1, 3, 5, ...) are uppercase

Example:
arrange("after be arrived two My so") returns "be ARRIVED two AFTER my SO"
*/

package kata

import "strings"

func Arrange(s string) string {
	// your solution here
}