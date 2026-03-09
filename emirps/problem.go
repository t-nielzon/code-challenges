/*
If you reverse the word "emirp" you will have the word "prime". That idea is
related with the purpose of this kata: we should select all the primes that
when reversed are a different prime (so palindromic primes should be discarded).

For example: 13, 17 are prime numbers and the reversed respectively are 31, 71
which are also primes, so 13 and 17 are "emirps". But primes 757, 787, 797 are
palindromic primes, meaning that the reversed number is the same as the
original, so they are not considered as "emirps" and should be discarded.

Your task: Create a function that receives one argument n, as an upper limit,
and returns [number_of_emirps_below_n, largest_emirp_below_n, sum_of_emirps_below_n]
*/
package kata

func FindEmirp(n int) [3]int {
}