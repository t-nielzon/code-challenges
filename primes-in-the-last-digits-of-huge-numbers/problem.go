// We have the following sequence:
//
//   f(0) = 0
//   f(1) = 1
//   f(2) = 1
//   f(3) = 2
//   f(4) = 4
//   f(n) = f(n-1) + f(n-2) - f(n-3) + f(n-4) - f(n-5)
//
// The first term of the sequence whose last nine digits form a prime number is
// 8480150779 (10 digits total), corresponding to the 92-th term, because its
// last nine digits, 480150779, is prime.
//
//   n-th term   k-th lastDig prime   term value          last9Digit   isPrime
//      92             1              8480150779           480150779      True
//      98             2              35922495169          922495169      True
//     110             3              644603021049         603021049      True
//     122             4              11566931883761       931883761      True
//     134             5              207560170886697      170886697      True
//
// Create a function KthLastDigPrime that receives k and returns the ordinal
// number of the term and the number formed by its last nine digits.
//
//   KthLastDigPrime(1) -> [92, 480150779]
//   KthLastDigPrime(2) -> [98, 922495169]
//   KthLastDigPrime(5) -> [134, 170886697]
//
// (Advice: use a fast primality test, such as Miller-Rabin.)

package kata

func KthLastDigPrime(k int) []int {
	return nil
}