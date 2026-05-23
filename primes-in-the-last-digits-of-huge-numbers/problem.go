package kata

/*
We have the following sequence:
f(0) = 0
f(1) = 1
f(2) = 1
f(3) = 2
f(4) = 4
f(n) = f(n-1) + f(n-2) - f(n-3) + f(n-4) - f(n-5)

The first term of the sequence that has its last nine digits forms a prime
number is the value 8480150779 (total of 10 digits), and corresponds to the
92-th term, because 480150779 is prime.

Create a function KthLastDigPrime(k) that receives the value of k as an
argument and outputs the ordinal number that corresponds to the term value
and the number formed by the last nine digits.

Examples:
    KthLastDigPrime(1) --> [92, 480150779]
    KthLastDigPrime(2) --> [98, 922495169]
    KthLastDigPrime(5) --> [134, 170886697]

Advise: Use a fast primality test, like Miller-Rabin.
*/

func KthLastDigPrime(k int) []int {
	// your code here
	return nil
}