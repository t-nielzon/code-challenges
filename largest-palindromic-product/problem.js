/*
 * A palindromic number is a number whose digits are in the same order when
 * read both forward and backward. For example, 11, 33, 1331, and 98389 are
 * all palindromic numbers.
 *
 * Write a function which, when given two integers `lower` and `upper`,
 * returns the largest palindromic number that's a product of any two integers
 * between the inclusive range of `lower` and `upper`.
 *
 * For example, if lower = 10 and upper = 99, the largest palindromic product
 * that can be made from two integers within the range is: 99 * 91 = 9009.
 *
 * You can use the same number twice, and the range is inclusive. So, for
 * lower = 1 and upper = 11, the largest palindrome would be 11 * 11 = 121.
 *
 * A single digit always counts as a palindrome.
 *
 * If no palindromic products can be made, return NaN.
 *
 * Arguments are always valid: upper > lower, and both are between 1 and 999999.
 *
 * This kata is mostly about optimization. Naïve solutions will time out.
 */

function largestPalindrome(lower, upper) {
  // your code here
}