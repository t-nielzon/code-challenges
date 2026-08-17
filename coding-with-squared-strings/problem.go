package main

/*
Coding with Squared Strings (5 kyu)

A squared string has n lines, each substring being n characters long.
For example: s = "abcd\nefgh\nijkl\nmnop" is a squared string of size 4.

We will use squared strings to code and decode texts. To make things easier
we suppose that our original text doesn't include the character '\n'.

CODING:
Input:
  - a text t of length l
  - Consider the smallest integer n such that n*n >= l
  - Complete t with the char of ascii code 11 until the length of t is n*n
  - Transform the new t in a squared string s of size n by inserting '\n' where needed
  - Apply a clockwise rotation of 90 degrees to s

DECODING:
Input:
  - a squared string s resulting from the coding
  - Apply a counter-clockwise rotation of 90 degrees to s
  - Remove the padding characters (ascii code 11) to get the original text

Example:
t = "I.was.going.fishing.that.morning.at.ten.o'clock"
code(t) -> "c.nhsoI\nltiahi.\noentinw\ncng.nga\nk..mg.s\n\voao.f.\n\v'trtig"
decode(code(t)) == "I.was.going.fishing.that.morning.at.ten.o'clock"
*/

func code(t string) string {
	// TODO: Implement
	return ""
}

func decode(s string) string {
	// TODO: Implement
	return ""
}