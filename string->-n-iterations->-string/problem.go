package main

/*
String -> N iterations -> String

We have a string s and a number n.

This kata applies a transformation n times to the string:
- Concatenate even-indexed characters (0, 2, 4, ...) to the front
- Concatenate odd-indexed characters (1, 3, 5, ...) to the back

Examples:
s = "Wow Example!"
result = "WwEapeo xml!"

s = "I'm Jomo Pipi"
result = "ImJm ii' ooPp"

For s = "qwertyuio" and n = 2:
after 1 iteration: s = "qetuowryi"
after 2 iterations: s = "qtorieuwy"
return "qtorieuwy"

Note: n can be greater than a billion, so optimization is required.
*/

func Repeat(s string, n int) string {
	return ""
}