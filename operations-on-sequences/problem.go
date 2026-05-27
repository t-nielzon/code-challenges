// Let an array or a list
//
//   arr = [x(1), x(2), x(3), x(4), ..., x(i), x(i+1), ..., x(m), x(m+1)]
//
// where all x(i) are positive integers.
// The length lg of this array will be a positive multiple of 4.
//
// Let
//
//   P = (x(1)**2 + x(2)**2) * (x(3)**2 + x(4)**2) * ... * (x(m)**2 + x(m+1)**2)
//
// x ** y means x raised to the power y.
//
// Task:
// Given an array or list arr the task is to find:
//   two nonnegative integers A and B such as P = A**2 + B**2  (1).
//
// The function Solve(arr) should return [A, B] where A and B verify (1).
//
// Examples:
//   Solve([2, 1, 3, 4]) returns [2, 11]:
//     (2*2 + 1*1) * (3*3 + 4*4) = 5 * 25 = 125 and 2*2 + 11*11 = 125
//
// Notes:
//   - The decomposition into A**2 + B**2 is not unique: the testing function
//     checks if (1) is verified.
//   - Lengths of lists are less than 100 and elements of lists less than 30.

package kata

import "math/big"

func Solve(arr []int) [2]*big.Int {
	return [2]*big.Int{}
}