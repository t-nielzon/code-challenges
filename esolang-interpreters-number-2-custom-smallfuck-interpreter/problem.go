/*
Esolang Interpreters #2 - Custom Smallfuck Interpreter

Smallfuck is an esoteric programming language which operates only on bits
with a limited data storage (tape).

Commands:
  > - Move pointer to the right (by 1 cell)
  < - Move pointer to the left (by 1 cell)
  * - Flip the bit at the current cell
  [ - Jump past matching ] if value at current cell is 0
  ] - Jump back to matching [ if value at current cell is nonzero

Termination conditions:
  - All commands have been considered from left to right
  - The pointer goes out-of-bounds

Implement interpreter(code, tape) that returns the final tape state as a string.
Non-command characters in code should be ignored.
*/
package kata

func Interpreter(code, tape string) string {
}