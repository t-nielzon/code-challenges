/*
You are given a string of n lines, each substring being n characters long.

We will study some transformations of this square of strings.

rot(s):
Clock rotation 180 degrees.
Example: rot("abcd\nefgh\nijkl\nmnop") => "ponm\nlkji\nhgfe\ndcba"

selfie_and_rot(s):
It is an initial string combined with its 180-degree clock-rotated version,
interspersed with dots proportional to the length of the segments.
Example: selfie_and_rot("abcd\nefgh\nijkl\nmnop") => 
"abcd....\nefgh....\nijkl....\nmnop....\n....ponm\n....lkji\n....hgfe\n....dcba"

oper(fct, s):
High-order function where fct is the function to apply to the string s.
fct will be one of rot or selfie_and_rot
*/

package main

func rot(s string) string {
}

func selfie_and_rot(s string) string {
}

func oper(fct func(string) string, s string) string {
}