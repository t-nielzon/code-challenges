// Emirps
//
// If you reverse the word "emirp" you will have the word "prime". That idea is
// related with the purpose of this kata: we should select all the primes that
// when reversed are a **different** prime (so palindromic primes should be
// discarded).
//
// For example: 13, 17 are prime numbers and the reversed respectively are 31,
// 71 which are also primes, so 13 and 17 are "emirps".
// But primes 757, 787, 797 are palindromic primes, meaning that the reversed
// number is the same as the original, so they are not considered as "emirps"
// and should be discarded.
//
// The emirps sequence is registered in OEIS as A006567.
//
// Your task:
// Create a function that receives one argument `n`, as an upper limit, and
// returns the following array:
//
// [number_of_emirps_below_n, largest_emirp_below_n, sum_of_emirps_below_n]
//
// Examples:
//   FindEmirp(10)  -> [0, 0, 0]      no emirps below 10
//   FindEmirp(50)  -> [4, 37, 98]    emirps below 50: 13, 17, 31, 37
//   FindEmirp(100) -> [8, 97, 418]   emirps below 100: 13,17,31,37,71,73,79,97

package kata

func FindEmirp(n int) [3]int {
	return [3]int{}
}