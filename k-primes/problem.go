/*
A natural number is called k-prime if it has exactly k prime factors, counted
with multiplicity. For example:

k = 2  -->  4, 6, 9, 10, 14, 15, 21, 22, ...
k = 3  -->  8, 12, 18, 20, 27, 28, 30, ...
k = 5  -->  32, 48, 72, 80, 108, 112, ...

A natural number is thus prime if and only if it is 1-prime.

Task:
Complete the function CountKprimes which is given parameters k, start, end
and returns a slice of the k-primes between start (inclusive) and end (inclusive).
For Go: nil slice is expected when there are no k-primes between start and end.

Second Task: puzzle
Given a positive integer s, find the total number of solutions of the equation
a + b + c = s, where a is 1-prime, b is 3-prime, and c is 7-prime.
*/
package kata

func CountKprimes(k, start, nd int) []int {
}

func Puzzle(s int) int {
}