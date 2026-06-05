// A Challenging Sequence
//
// Each composed integer may have a prime factorization. If n is the integer,
// it will be n = p1^k1 . p2^k2 . p3^k3 ..... pn^kn, where k1, k2, k3...., kn
// are the exponents corresponding to each prime in the factorization.
//
// The radical of n, Rad(n), will be the product of all the prime factors of
// the number (without the exponents), so:
//
//   Rad(n) = p1 . p2 . p3 ..... pn
//
// Example: if n = 172800 its prime factorization is n = 2^8 . 3^3 . 5^2, so
// Rad(172800) = 2 . 3 . 5 = 30.
//
// Many numbers smaller than 172800 may have a radical considerably higher,
// e.g. Rad(9305) = 5 . 1861 = 9305.
//
// Primes are a particular case, the value of its radical coincides with the
// value of the prime itself: Rad(31) = 31, Rad(57) = 57.
// To number 1 it's assigned Rad(1) = 1.
//
// Comparison for the first 10 numbers:
//
//   n       Prime Factors     Radical(n)
//   1                                  1
//   2          2                       2
//   3          3                       3
//   4          2                       2
//   5          5                       5
//   6          2, 3                    6
//   7          7                       7
//   8          2                       2
//   9          3                       3
//  10          2, 5                   10
//
// The list may be sorted by the value of the radicals (ties broken by the
// number n itself), producing another sequence with an ordinal number k:
//
//  k- Term     n     Radical(n)
//  1           1            1
//  2           2            2
//  3           4            2
//  4           8            2
//  5           3            3
//  6           9            3
//  7           5            5
//  8           6            6
//  9           7            7
//  10         10           10
//
// Create a function hashRadSeq() that receives two arguments:
//
//   - nMax: upper bound such that all values of n are in 1 <= n <= nMax
//   - k:    the ordinal number of a certain term in the sequence sorted by
//           the values of Rad(n)
//
// The function should output the value of n for the corresponding ordinal
// number k (k-th), in a sequence sorted by the values of Rad(n).
//
// Examples:
//   hashRadSeq(10, 4) ------> 8
//   hashRadSeq(10, 6) ------> 9
//   hashRadSeq(10, 9) ------> 7
//   hashRadSeq(20, 4) ------> 8
//   hashRadSeq(20, 6) ------> 3
//   hashRadSeq(20, 9) ------> 6
//
// Tested for nMax up to 60000 and k up to 11000.

package kata

func hashRadSeq(nMax int, k int) int {
	return 0
}