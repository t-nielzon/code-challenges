/*
Primes in the Last Digits of Huge Numbers

We have the following sequence:
  f(0) = 0
  f(1) = 1
  f(2) = 1
  f(3) = 2
  f(4) = 4
  f(n) = f(n-1) + f(n-2) - f(n-3) + f(n-4) - f(n-5)

The first term of the sequence that has its last nine digits forming a prime number
is 8480150779 (at n=92), where 480150779 is prime.

Create a function kthLastDigPrime(k) that returns the ordinal number (n) and the last 9 digits
for the k-th term where the last 9 digits form a prime number.

Example:
  kthLastDigPrime(1) => [92, 480150779]
  kthLastDigPrime(2) => [98, 922495169]
  kthLastDigPrime(5) => [134, 170886697]

Use a fast primality test such as Miller-Rabin.
*/
package main

func kthLastDigPrime(k int) [2]int {
	// TODO: implement
	return [2]int{}
}