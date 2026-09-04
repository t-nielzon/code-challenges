package main

/*
A Challenging Sequence

Each composed integer may have a prime factorization.
The radical of n, Rad(n), is the product of all the prime factors of the number (without the exponents).

For example, if n = 172800:
  n = 2^8 . 3^3 . 5^2
  Rad(172800) = 2 . 3 . 5 = 30

Create a function HashRadSeq() that receives two arguments:
  - nMax: the upper bound such that all values of n are in the interval (1, nMax) (1 <= n <= nMax)
  - k: the ordinal number of a term in a sequence sorted by the values of the radicals

The function should output the value of n for the k-th position in a sequence sorted by Rad(n).
When multiple numbers have the same radical, they are ordered by their own value.

Example:
  HashRadSeq(10, 4) -> 8
  HashRadSeq(10, 6) -> 9
  HashRadSeq(10, 9) -> 7
*/

func HashRadSeq(nMax, k int) int {

}