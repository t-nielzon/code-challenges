package main

/*
Emirp kata

If you reverse the word "emirp" you will have the word "prime".
An emirp is a prime number that when reversed is a different prime.
Palindromic primes should be discarded.

Example: 13, 17 are primes and reversed are 31, 71 which are also primes,
so 13 and 17 are emirps.

Your task: create a function that receives one argument n as an upper limit,
and return [number_of_emirps_below_n, largest_emirp_below_n, sum_of_emirps_below_n]

Examples:
find_emirp(10) -> [0, 0, 0]
find_emirp(50) -> [4, 37, 98]
find_emirp(100) -> [8, 97, 418]

Advice: Use a prime generator (Sieve of Eratosthenes) and store primes in a set
for faster lookup rather than testing primality for each number.
*/

func FindEmirp(n int) [3]int {
	// TODO: implement
	return [3]int{0, 0, 0}
}