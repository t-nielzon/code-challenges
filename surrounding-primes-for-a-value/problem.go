// We need a function primeBefAft() that gives the largest prime below a
// certain given value n (befPrime), and the smallest prime larger than this
// value (aftPrime).
//
// The result should be output in a list like the following:
//
//   primeBefAft(n) == [befPrime, aftPrime]
//
// If n is a prime number it will give two primes; n will not be included
// in the result.
//
// Examples:
//
//   primeBefAft(100) == [97, 101]
//   primeBefAft(97)  == [89, 101]
//   primeBefAft(101) == [97, 103]
//
// Range for the random tests: 1000 <= n <= 200000
// (The extreme and special case n = 2 will not be considered for the tests.)

package kata

func PrimeBefAft(n int) [2]int {
	return [2]int{}
}