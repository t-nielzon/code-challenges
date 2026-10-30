package main

/*
Each composed integer may have a prime factorization.
The radical of n, Rad(n), will be the product of all the prime factors of the number, (without the exponents).

For example: n = 172800 = 2^8 . 3^3 . 5^2
Rad(172800) = 2 . 3 . 5 = 30

Create a function hashRadSeq() that receives two arguments:
- nMax is the upper bound such that all values of n are in the interval [1, nMax]
- k is the ordinal number of a certain term in a sequence sorted by the values of Rad(n)

The function should output the value of n for the corresponding ordinal number k in a sequence sorted by Rad(n).
Numbers with the same radical are ordered by their own value.

Examples:
hashRadSeq(10, 4) -> 8
hashRadSeq(10, 6) -> 9
hashRadSeq(20, 4) -> 8
*/

func hashRadSeq(nMax int, k int) int {
	return 0
}