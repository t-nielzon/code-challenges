package solution

/*
Iterative Rotation Cipher (IRC)

Implement two functions, Encode and Decode.

Encode receives a positive integer n and a string value.
Decode receives one string value.
Each function returns a string value.

How It Works
Encoding and decoding are done by performing a series of character and
substring rotations on a string input.

Encoding: The number of rotations is determined by the value of n. The
sequence of rotations is applied in the following order:
  Step 1: remove all spaces in the string (but remember their positions)
  Step 2: shift the order of characters in the new string to the right by n
  Step 3: put the spaces back in their original positions
  Step 4: shift the characters of each substring (separated by one or more
          consecutive spaces) to the right by n

Repeat this process until it has been completed n times in total.
The value n is then prepended to the resulting string with a space.

Decoding: Decoding simply reverses the encoding process.

Technical Details
- Input will always be valid.
- The characters used in the strings include any combination of alphanumeric
  characters, the space character, the newline character, and any of the
  following: _!@#$%^&()[]{}+-*"/="'<>,.?:;
*/

func Encode(n int, s string) string {
	panic("not implemented")
}

func Decode(s string) string {
	panic("not implemented")
}