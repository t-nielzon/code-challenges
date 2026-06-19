// Odd-Even Compositions
// Difficulty: 5 kyu
//
// A composition of a number N is a way of writing N as an ordered sum of
// natural numbers (order matters). The number 3 has 4 compositions:
//   3 / 2+1 / 1+2 / 1+1+1
//
// An "odd-even composition" restricts which parts are allowed. Reading a
// composition left to right and subtracting each part from a running value
// (which starts at N), a part k is valid only if it is:
//   - 1                                  (always valid),
//   - equal to the current running value (always valid), or
//   - of the same parity as the current running value.
// A composition is an odd-even composition only if every one of its parts
// is valid.
//
// For example with N = 4:
//   "3+1"   is invalid (3 is odd, 4 is even)
//   "1+2+1" is invalid (after the 1, running value is 3; 2 is even)
//   "2+2"   is valid
//   "1+3"   is valid
// N = 4 has 6 odd-even compositions, N = 3 has 3.
//
// Return the number of odd-even compositions of the given number.
// Inputs range from 0 to 10^3 inclusive. If input is 0 return 1.
//
//   OddEvenCompositions(0)  -> 1
//   OddEvenCompositions(7)  -> 27
//   OddEvenCompositions(28) -> 3188646

package main

import "math/big"

func OddEvenCompositions(n int) *big.Int {
	// your code here
	return nil
}