/*
A Challenging Sequence

Each composed integer may have a prime factorization.
If n is the integer, it will be:

n = p1^k1 . p2^k2 . p3^k3 ..... pn^kn

and k1, k2, k3....., kn are the exponents corresponding to each prime in the factorization.

The radical of n, Rad(n), will be the product of all the prime factors of the number,
(without the exponents), so:

Rad(n) = p1 . p2 . p3 .....pn

Let's see an example for a number, if n = 172800
its prime factorization will be
n = 2^8 . 3^3 . 5^2,
Rad(172800) = 2 . 3 . 5 = 30

Rad(1) = 1

Create a function HashRadSeq() that receives two arguments:
- nMax is the upper bound such that all values of n are in the interval (1, nMax) (1 <= n <= nMax)
- the number k, which is the ordinal number of a certain term in a sequence sorted by the values of the radicals
- the function should output the value of n, for the corresponding ordinal number k (k-th),
  in a sequence sorted by the values of Rad(n)

Examples:
HashRadSeq(10, 4) -> 8
HashRadSeq(10, 6) -> 9
HashRadSeq(10, 9) -> 7
*/
package kata

func HashRadSeq(nMax, k int) int {
}