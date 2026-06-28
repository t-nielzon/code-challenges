/*
Steps in Primes
Difficulty: 6 kyu

The prime numbers are not regularly spaced. For example from 2 to 3 the step is 1.
From 3 to 5 the step is 2. From 7 to 11 it is 4.
Between 2 and 50 we have the following pairs of 2-steps primes:
3, 5 - 5, 7, - 11, 13, - 17, 19, - 29, 31, - 41, 43

Write a function step with parameters:
- g (integer >= 2): the step we are looking for
- m (integer >= 2): start of the search (inclusive)
- n (integer >= m): end of the search (inclusive)

Return the first pair of prime numbers spaced with a step of g between m and n,
or nil if no such pair exists.

Examples:
step(2, 5, 7)    --> []int{5, 7}
step(2, 5, 5)    --> nil
step(4, 130, 200) --> []int{163, 167}
step(2, 4900, 4919) --> nil
*/

package kata

func Step(g, m, n int) []int {
	// stub
	return nil
}