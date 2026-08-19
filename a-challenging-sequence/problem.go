package kata

/*
Each composed integer may have a prime factorization.
If n is the integer, its prime factorization will be:
n = p1^k1 * p2^k2 * p3^k3 * ... * pn^kn

The radical of n, Rad(n), is the product of all distinct prime factors
(without the exponents):
Rad(n) = p1 * p2 * p3 * ... * pn

For example:
- n = 172800 = 2^8 * 3^3 * 5^2
- Rad(172800) = 2 * 3 * 5 = 30
- Rad(31) = 31 (prime)
- Rad(1) = 1

Create a function that:
1. Computes the radical for each number from 1 to nMax
2. Sorts numbers by their radical value (and by n itself when radicals are equal)
3. Returns the n-value at the k-th position in the sorted sequence

Examples:
hash_radSeq(10, 4) -> 8
hash_radSeq(10, 6) -> 9
hash_radSeq(10, 9) -> 7
hash_radSeq(20, 4) -> 8
hash_radSeq(20, 6) -> 3
hash_radSeq(20, 9) -> 6
*/

func HashRadSeq(nMax int, k int) int {
	return 0
}