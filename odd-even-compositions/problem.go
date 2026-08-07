package main

/*
# Odd-Even Compositions

Return the number of odd-even compositions by which the given number can be expressed.

An odd-even composition is a composition where each element must:
- Be 1 (always valid), or
- Share the same parity (odd/even) as the current remaining value

For example, N=4 (even):
- 4: valid
- 2+2: valid (2 is even, matches parity of 4; then 2 matches parity of 2)
- 2+1+1: valid
- 1+3: valid (1 is valid; 3 matches parity of remaining 3)
- 1+1+2: valid
- 1+1+1+1: valid
- Total: 6 valid compositions

For example, N=3 (odd):
- 3: valid
- 1+2: valid (1 is valid; 2 matches parity of remaining 2)
- 1+1+1: valid
- Total: 3 valid compositions

Examples:
  OddEvenCompositions(0)  // returns 1
  OddEvenCompositions(3)  // returns 3
  OddEvenCompositions(4)  // returns 6
  OddEvenCompositions(7)  // returns 27
  OddEvenCompositions(28) // returns 3188646
*/

func OddEvenCompositions(n int) int {
	return 0
}