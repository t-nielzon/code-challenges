package kata

/*
You are given a string of n lines, each substring being n characters long.

Transformations:

rot(s): Clock rotation 180 degrees.
Example: "abcd\nefgh\nijkl\nmnop" => "ponm\nlkji\nhgfe\ndcba"

selfie_and_rot(s): Initial string combined with its 180-degree rotated version,
interspersed with dots proportional to the length of segments.
Example: "abcd\nefgh\nijkl\nmnop" =>
"abcd....\nefgh....\nijkl....\nmnop....\n....ponm\n....lkji\n....hgfe\n....dcba"

oper(fct, s): High-order function where fct is the function to apply to string s.
fct will be one of: rot, selfie_and_rot
*/

import "strings"

type Operation func(string) string

func rot(s string) string {
	// 180-degree rotation of square string
}

func selfie_and_rot(s string) string {
	// Original string combined with 180-degree rotated version with dots
}

func oper(fct Operation, s string) string {
	// Apply transformation function fct to string s
}