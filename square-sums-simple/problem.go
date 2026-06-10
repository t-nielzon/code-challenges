// # Square sums
//
// Write function that, given integer number `N` (in range `2..43`), returns
// array of integers `1..N` arranged in a way, so sum of each 2 consecutive
// numbers is a square.
//
// Solution is valid if and only if following two criterias are met:
//   1. Each number in range `1..N` is used once and only once.
//   2. Sum of each 2 consecutive numbers is a perfect square.
//
// ### Example
//
// For N=15 solution could look like this:
//   [ 9, 7, 2, 14, 11, 5, 4, 12, 13, 3, 6, 10, 15, 1, 8 ]
//
// Sum of each 2 consecutive numbers is a perfect square (9, 16 or 25).
//
// If there is no solution, return false (an empty/nil slice). For example if
// N=5, then numbers 1,2,3,4,5 cannot be put into a square sums row.
//
// Returns the arrangement as []int, or nil if no solution exists.
package kata

func SquareSumsRow(n int) []int {
	return nil
}