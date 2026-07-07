package main

// Moves in squared strings (II)
//
// You are given a string of n lines, each substring being n characters long.
// Example: s = "abcd\nefgh\nijkl\nmnop"
//
// Rot(s): Clock rotation 180 degrees.
// Example: Rot("abcd\nefgh\nijkl\nmnop") => "ponm\nlkji\nhgfe\ndcba"
//
// SelfieAndRot(s): Initial string combined with its 180-degree rotated version,
// interspersed with dots proportional to the length of the line segments.
// Example: SelfieAndRot("abcd\nefgh\nijkl\nmnop") => "abcd....\nefgh....\nijkl....\nmnop....\n....ponm\n....lkji\n....hgfe\n....dcba"
//
// Oper(fct, s): Higher-order function where fct is a function to apply to string s.
// fct will be one of Rot or SelfieAndRot.

func Rot(s string) string {
}

func SelfieAndRot(s string) string {
}

func Oper(fct func(string) string, s string) string {
}