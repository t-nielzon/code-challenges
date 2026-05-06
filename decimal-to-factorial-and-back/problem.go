package kata

/*
Coding decimal numbers with factorials is a way of writing out numbers
in a base system that depends on factorials, rather than powers of numbers.

In this system, the last digit is always 0 and is in base 0!. The digit before
that is either 0 or 1 and is in base 1!. The digit before that is either
0, 1, or 2 and is in base 2!, etc.
More generally, the nth-to-last digit is always 0, 1, 2, ..., n and is in base n!.

Example:
The decimal number 463 is encoded as "341010", because:
463 = 3×5! + 4×4! + 1×3! + 0×2! + 1×1! + 0×0!

If we are limited to digits 0..9, the biggest number we can encode is 10!-1.
So we extend 0..9 with letters A..Z. With these 36 digits we can now encode
numbers up to 36!-1.

Task:
Implement two functions:
  - Dec2FactString: receives a decimal number, returns its factorial representation.
  - FactString2Dec: receives a factorial representation string, returns the decimal number.

Given numbers will always be positive.
*/

import "math/big"

func Dec2FactString(nb *big.Int) string {
	// TODO: implement
	return ""
}

func FactString2Dec(s string) *big.Int {
	// TODO: implement
	return nil
}