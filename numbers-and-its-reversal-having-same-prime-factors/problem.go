// problem.go
package main

/*
The number 1089 is the smallest one, non palindromic, that has the same prime factors that its reversal.
Thus:
  prime factorization of 1089 with 3, 3, 11, 11 -------> 3, 11
  prime factorization of 9801 with 3, 3, 3, 3, 11, 11 -------> 3, 11

The task is to create a function SameFactRev() that receives a nMax, to find all the numbers
with the above property, below nMax.

The function SameFactRev() will output a sorted list with the found numbers below nMax.

Examples:
  SameFactRev(1100) -----> [1089]
  SameFactRev(2500) -----> [1089, 2178]

Palindromic numbers should be discarded.
*/

func SameFactRev(nMax int) []int {
	// TODO: implement
	return []int{}
}