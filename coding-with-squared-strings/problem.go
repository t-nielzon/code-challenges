package kata

/*
Coding with Squared Strings

A squared string has n lines, each substring being n characters long.
Example: s = "abcd\nefgh\nijkl\nmnop" is a squared string of size 4.

Coding:
- Input: text t of length l.
- Find smallest integer n such that n*n >= l.
- Pad t with char of ascii code 11 until length is n*n.
- Transform t into a squared string s of size n by inserting '\n' where needed.
- Apply a clockwise rotation of 90 degrees to s.

Decoding:
- Input: a squared string s from coding.
- Apply a counter-clockwise rotation of 90 degrees.
- Clean to recover original text t.
*/

func Code(s string) string {
}

func Decode(s string) string {
}