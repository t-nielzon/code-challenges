/*
Simple assembler interpreter

We want to create a simple interpreter of assembler which will support the following instructions:

- mov x y - copies y (either a constant value or the content of a register) into register x
- inc x - increases the content of the register x by one
- dec x - decreases the content of the register x by one
- jnz x y - jumps to an instruction y steps away (positive means forward, negative means backward,
  y can be a register or a constant), but only if x (a constant or a register) is not zero

Register names are alphabetical (letters only). Constants are always integers (positive or negative).

The function will take an input list with the sequence of the program instructions and will execute them.
The program ends when there are no more instructions to execute, then it returns a map with the contents
of the registers.
*/
package kata

func SimpleAssembler(program []string) map[string]int {
}