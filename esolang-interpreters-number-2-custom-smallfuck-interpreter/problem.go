package main

/*
Esolang Interpreters #2 - Custom Smallfuck Interpreter

Smallfuck is an esoteric programming language operating on bits.

Commands:
- > : Move pointer right (terminate if out of bounds)
- < : Move pointer left (terminate if out of bounds)
- * : Flip the bit at current cell
- [ : Jump past matching ] if current cell is 0
- ] : Jump back to matching [ if current cell is nonzero

The interpreter receives:
- code: The Smallfuck program as a string (may contain non-command characters)
- tape: The initial tape state as a string of 0s and 1s

Returns:
- The final state of the tape as a string

The pointer starts at the leftmost cell. Execution terminates when all commands
are processed or when the pointer moves out of bounds.
*/

func Interpreter(code string, tape string) string {
	return ""
}