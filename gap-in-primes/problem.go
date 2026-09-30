package main

/*
The prime numbers are not regularly spaced. For example from 2 to 3 the gap is 1.
From 3 to 5 the gap is 2. From 7 to 11 it is 4.

A prime gap of length n is a run of n-1 consecutive composite numbers between
two successive primes (http://mathworld.wolfram.com/PrimeGaps.html).

Parameters:
- g (integer >= 2): the gap we are looking for
- m (integer > 2): start of search (inclusive)
- n (integer >= m): end of search (inclusive)

Return the first pair of two successive prime numbers with a gap of g between
the limits m and n. Return nil if no such pair exists.

Examples:
- gap(2, 5, 7) --> [5, 7]
- gap(2, 5, 5) --> nil
- gap(4, 130, 200) --> [163, 167]
- gap(6, 100, 110) --> nil
*/

func gap(g int, m int, n int) []int {
	return nil
}