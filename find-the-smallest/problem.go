// You have a positive number n consisting of digits.
// You can do at most one operation:
// Choosing the index of a digit in the number, remove this digit at that
// index and insert it back to another or at the same place in the number
// in order to find the smallest number you can get.
//
// Task:
// Return an array [3]int64 with
//   1) the smallest number you got
//   2) the index i of the digit d you took, i as small as possible
//   3) the index j (as small as possible) where you insert this digit d
//
// Examples:
//   Smallest(261235) --> [126235, 2, 0]
//   Smallest(209917) --> [29917, 0, 1]
//   Smallest(1000000) --> [1, 0, 6]

package kata

func Smallest(n int64) [3]int64 {
	return [3]int64{}
}