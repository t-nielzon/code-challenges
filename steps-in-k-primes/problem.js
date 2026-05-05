/*
 * A natural number is called k-prime if it has exactly k prime factors, counted with multiplicity.
 *
 * A natural number is thus prime if and only if it is 1-prime.
 *
 * Examples of k-primes:
 * k = 2 -> 4, 6, 9, 10, 14, 15, 21, 22, …
 * k = 3 -> 8, 12, 18, 20, 27, 28, 30, …
 * k = 5 -> 32, 48, 72, 80, 108, 112, …
 *
 * Task:
 * Write a function kprimes_step(k, step, start, nd) with parameters:
 * - k (integer > 0): indicates the type of k-primes we are looking for
 * - step (integer > 0): indicates the step we want to find between two k-primes
 * - start (integer >= 0): gives the start of the search (start inclusive)
 * - nd (integer >= start): gives the end of the search (nd inclusive)
 *
 * Returns an array of all the pairs of k-prime numbers spaced with a step of step
 * between the limits start, nd. This array can be empty.
 *
 * Examples:
 * kprimes_step(2, 2, 0, 50) => [[4, 6], [33, 35]]
 * kprimes_step(6, 14, 2113665, 2113889) => [[2113722, 2113736]]
 * kprimes_step(2, 10, 0, 50) => [[4, 14], [15, 25], [25, 35], [39, 49]]
 * kprimes_step(5, 20, 0, 50) => []
 */

function kprimes_step(k, step, start, nd) {
    
}