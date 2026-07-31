/*
 * Buddy pairs
 * 
 * The divisors of a positive integer n are said to be proper when considering
 * only the divisors other than n itself.
 * 
 * Let s(n) be the sum of these proper divisors of n.
 * 
 * Call buddy two positive integers such that the sum of the proper divisors
 * of each number is one more than the other number:
 * (n, m) are a pair of buddy if s(m) = n + 1 and s(n) = m + 1
 * 
 * For example 48 & 75 is such a pair:
 * - Divisors of 48 are: 1, 2, 3, 4, 6, 8, 12, 16, 24 --> sum: 76 = 75 + 1
 * - Divisors of 75 are: 1, 3, 5, 15, 25 --> sum: 49 = 48 + 1
 * 
 * Task: Given two positive integers start and limit, the function buddy(start, limit)
 * should return the first pair (n m) of buddy pairs such that n is between start
 * (inclusive) and limit (inclusive); m can be greater than limit and has to be
 * greater than n.
 * 
 * If there is no buddy pair satisfying the conditions, return "Nothing".
 */

function buddy(start, limit) {
  
}