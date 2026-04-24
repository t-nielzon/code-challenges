/*
Odd-Even Compositions

Given a natural number N, count the number of ways to express N as an ordered
sum of natural numbers (a composition) subject to the following rules applied
left-to-right while a running remainder R (initially N) is decreased:

 - A part equal to 1 is always valid (R becomes R-1).
 - A part equal to R itself is always valid (R becomes 0).
 - Any other part k (2 <= k < R) is valid only if k has the same parity as R.

Return the total number of such compositions.

Notes:
 - If input is 0, return 1.
 - N ranges from 0 to 10^3 inclusive, so the result may be very large.

Examples:
 OddEvenCompositions(0)  == 1
 OddEvenCompositions(7)  == 27
 OddEvenCompositions(28) == 3188646
*/

package kata

import "math/big"

func OddEvenCompositions(n int) *big.Int {
	// TODO: implement
	return nil
}