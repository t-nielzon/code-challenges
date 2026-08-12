/*
Odd-Even Compositions

A composition is an ordered way to express a number as a sum of natural numbers.
An odd-even composition restricts which numbers can be used at each step.

At each step when we have a remaining sum R, we can use:
- 1 (always valid)
- Any number with the same parity (odd/even) as R (where 1 < k < R)
- R itself (completes the composition)

For N=4 (even):
- 4: valid (use 4 itself)
- 3+1: invalid (3 is odd, 4 is even - different parity)
- 2+2: valid (2 is even like 4, then 2 is itself)
- 2+1+1: valid (2 is even like 4, 1 always valid, 1 is itself)
- 1+3: valid (1 always valid, 3 is itself)
- 1+2+1: invalid (after 1: remaining=3 (odd); 2 is even - different parity)
- 1+1+2: valid (1 always valid, 1 always valid, 2 is itself)
- 1+1+1+1: valid (all 1s always valid)
Total: 6 valid compositions

Task: Return the number of odd-even compositions of N.
If N=0 return 1.

Examples:
- OddEvenCompositions(0) returns 1
- OddEvenCompositions(7) returns 27
- OddEvenCompositions(28) returns 3188646
*/

package kata

func OddEvenCompositions(n int) int {
	return 0
}