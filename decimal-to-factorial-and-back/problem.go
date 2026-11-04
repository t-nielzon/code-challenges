package main

/*
Coding decimal numbers with factorials is a way of writing out numbers
in a base system that depends on factorials, rather than powers of numbers.

In this system, the last digit is always `0` and is in base 0!. The digit before that is either `0 or 1` and is in base 1!.  The digit before that is either `0, 1, or 2` and is in base 2!, etc.
More generally, the nth-to-last digit is always `0, 1, 2, ..., n` and is in base n!.

Example: The decimal number 463 is encoded as "341010", because:
463 = 3×5! + 4×4! + 1×3! + 0×2! + 1×1! + 0×0!

Digits 0-9 represent values 0-9, and A-Z represent values 10-35.

Task: Convert between decimal and factorial number system representation.
- ToFactorial: convert decimal to factorial representation string
- FromFactorial: convert factorial representation string to decimal
*/

func ToFactorial(decimal int) string {
	// TODO: implement
	return ""
}

func FromFactorial(factorial string) int {
	// TODO: implement
	return 0
}