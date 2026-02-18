/*
Odd-Even Compositions

Given a number N, return the number of "odd-even compositions" by which it can be expressed.

In a composition, N is expressed as an ordered sum of natural numbers.
In an odd-even composition, each part (other than 1 or the remaining number itself)
must share the same parity (odd/even) as the current remainder.

Rules:
- 1 is always valid (gray)
- The number itself (remainder) is always valid (gray)
- A part that shares parity with the current remainder is valid (green)
- A part that does NOT share parity with the current remainder is invalid (red)
- If any part is invalid, the entire composition is invalid

If input is 0, return 1.
Numbers range from 0 to 10^3 inclusive; solution must be efficient.

Examples:
  OddEvenCompositions(0)  => 1
  OddEvenCompositions(3)  => 3
  OddEvenCompositions(4)  => 6
  OddEvenCompositions(7)  => 27
  OddEvenCompositions(28) => 3188646
*/
package kata

import "math/big"

func OddEvenCompositions(n int) *big.Int {

}