/*
## Task

Create a function LongestComb that takes a sequence of integers and a method of
ordering integers (increasing/decreasing) and returns all of the longest
subsequences (with length of at least 3) of the given sequence such that the
subsequence elements are ordered according to the given ordering method.

A subsequence is a sequence derived by deleting some or no elements from a given
sequence without changing the order of the remaining elements.

#### Arguments
- arr: An array of integers
- command: A string representing the ordering ("< <" or "<<" for increasing,
  "> >" or ">>" for decreasing)

#### Output
- A 2D slice of the longest subsequences (length >= 3) ordered by original
  element indices from arr.
- If no valid subsequence exists, return an empty slice.

#### Examples
  [-1, 3, -34, 18, -55, 60, 118, -64], "< <" --> [[-1, 3, 18, 60, 118]]
  [-1, 3, -34, 18, -55, 60, 118, -64], "> >" --> [[-1, -34, -55, -64], [3, -34, -55, -64]]
*/
package kata

func LongestComb(arr []int, command string) [][]int {
}