package main

/*
Primes in the Last Digits of Huge Numbers

We have the following sequence:
  f(0) = 0
  f(1) = 1
  f(2) = 1
  f(3) = 2
  f(4) = 4
  f(n) = f(n-1) + f(n-2) - f(n-3) + f(n-4) - f(n-5)

The first term of the sequence that has its last nine digits forming a prime
number is the value 8480150779 (total of 10 digits), corresponding to the
92-nd term, because 480150779 is prime.

Create a function kthLastDigPrime() that receives the value of k as an
argument and returns the ordinal number that corresponds to the term and
the number formed by the last nine digits.

Examples:
  kthLastDigPrime(1) => [92, 480150779]
  kthLastDigPrime(2) => [98, 922495169]
  kthLastDigPrime(5) => [134, 170886697]

Advice: Use a fast primality test such as Miller-Rabin.
*/

func kthLastDigPrime(k int) []int64 {
	return []int64{}
}