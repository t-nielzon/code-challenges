package main

/*
The prime numbers are not regularly spaced. For example from 2 to 3 the gap is 1.
From 3 to 5 the gap is 2. From 7 to 11 it is 4.

A prime gap of length n is a run of n-1 consecutive composite numbers between two successive primes.

We need to find the first pair of successive primes with a gap of g between m and n (inclusive).

Parameters:
- g (integer >= 2): the gap we are looking for
- m (integer > 2): start of the search (m inclusive)
- n (integer >= m): end of the search (n inclusive)

Returns:
- [p1, p2] where p1 and p2 are consecutive primes with gap g
- nil if no such pair exists

Examples:
- gap(2, 5, 7) --> [5, 7]
- gap(2, 5, 5) --> nil
- gap(4, 130, 200) --> [163, 167]
*/

func Gap(g int, m int, n int) []int {
	// function body goes here
}