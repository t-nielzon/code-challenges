/*
Second Variation on Caesar Cipher

In this country soldiers are poor but they need a certain level of secrecy
for their communications so they reinvent Caesar cipher in the following way.

They shift each letter by a given amount (cyclic a-z, A-Z). Non-letter
characters are kept as-is.

They prefix the coded message with 2 letters: the first letter of the
original message (downcased), and that letter shifted by the rotate amount.

They then split the result into 5 parts where parts 1-4 are equal length
and part 5 is shorter or equal. If the fifth part is empty, omit it.

Encode: given a string and shift, return the array of parts.
Decode: given the array of parts, return the original string.
*/
package kata

func Encode(s string, shift int) []string {
	return nil
}

func Decode(arr []string) string {
	return ""
}