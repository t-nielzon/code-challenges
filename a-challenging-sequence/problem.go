/*
Each composed integer may have a prime factorization.
If n is the integer, it will be:

n = p1^k1 * p2^k2 * ... * pn^kn

and k1, k2, k3....., kn are the exponents corresponding to each prime in the factorization.

The radical of n, Rad(n), will be the product of all the prime factors of the number,
(without the exponents), so:

Rad(n) = p1 * p2 * p3 * ... * pn

For example, if n = 172800, its prime factorization is 2^8 * 3^3 * 5^2,
so Rad(172800) = 2 * 3 * 5 = 30.

Primes have Rad(p) = p, and Rad(1) = 1.

Create hashRadSeq(nMax, k) that returns the value of n at the k-th position
in the sequence of integers 1..nMax sorted by Rad(n) ascending, ties broken
by n ascending.

Examples:
  hashRadSeq(10, 4) -> 8
  hashRadSeq(10, 6) -> 9
  hashRadSeq(10, 9) -> 7
  hashRadSeq(20, 4) -> 8
  hashRadSeq(20, 6) -> 3
  hashRadSeq(20, 9) -> 6

nMax up to 60000, k up to 11000.
*/

package kata

func HashRadSeq(nMax, k int) int {

}