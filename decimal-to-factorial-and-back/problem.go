/*
Coding decimal numbers with factorials is a way of writing out numbers
in a base system that depends on factorials, rather than powers of numbers.

In this system, the last digit is always 0 and is in base 0!. The digit before
that is either 0 or 1 and is in base 1!. The digit before that is either 0, 1,
or 2 and is in base 2!, etc. More generally, the nth-to-last digit is always
0, 1, 2, ..., n and is in base n!.

The decimal number 463 is encoded as "341010", because:
463 = 3×5! + 4×4! + 1×3! + 0×2! + 1×1! + 0×0!

If we are limited to digits 0..9, the biggest number we can encode is 10!-1
(= 3628799). So we extend 0..9 with letters A..Z. With these 36 digits we can
now encode numbers up to 36!-1.

Task: two functions.
- DecToFact: receive a decimal number, return a string with the factorial representation.
- FactToDec: receive a factorial representation string, produce the decimal number.
*/
package kata

import "math/big"

func DecToFact(nb int64) string {
	return ""
}

func FactToDec(s string) int64 {
	return 0
}