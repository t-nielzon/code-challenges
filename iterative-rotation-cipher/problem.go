/*
Iterative Rotation Cipher (5 kyu)

Implement an iterative rotation cipher with encode and decode functions.

Encoding with parameter n and a text string:
1. Remove all spaces (remember their positions)
2. Rotate the characters right by n
3. Put spaces back in original positions
4. Rotate each space-separated substring right by n
5. Repeat steps 1-4 a total of n times
6. Prepend n and a space to the result

Decoding reverses the encoding process.

Example:
  quote := "If you wish to make an apple pie from scratch, you must first invent the universe."
  Encode(10, quote) returns "10 hu fmo a,ys vi utie mr snehn rni tvte .ysushou teI fwea pmapi apfrok rei tnocsclet"
*/
package main

func Encode(n int, text string) string {
	// TODO: Implement encoding logic
	return ""
}

func Decode(text string) string {
	// TODO: Implement decoding logic
	return ""
}