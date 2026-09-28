package main

/*
A Rule of Divisibility by 13

A divisibility rule is a shorthand way of determining whether a given integer is divisible by a fixed divisor
without performing the division, usually by examining its digits.

When you divide the successive powers of 10 by 13 you get the following remainders:
1, 10, 9, 12, 3, 4 (and the pattern repeats)

The method:
- Multiply the rightmost digit by 1, the second rightmost by 10, etc. (following the sequence)
- Sum all these products
- Repeat until the result is stationary (no longer changes)
- Return the final stationary number

Example: For 1234567
- 7×1 + 6×10 + 5×9 + 4×12 + 3×3 + 2×4 + 1×1 = 178
- 8×1 + 7×10 + 1×9 = 87
- 7×1 + 8×10 = 87
- Returns 87 (stationary)
*/

func DivisibilityRule13(n int64) int64 {
	// TODO
}