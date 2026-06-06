/*
 * Closest and Smallest
 *
 * Input:
 *   a string `strng` of n positive numbers (n = 0 or n >= 2)
 *
 * Let us call weight of a number the sum of its digits.
 * For example 99 will have "weight" 18, 100 will have "weight" 1.
 * Two numbers are "close" if the difference of their weights is small.
 *
 * Task:
 *   For each number in `strng` calculate its "weight" and then find two
 *   numbers of `strng` that have:
 *     - the smallest difference of weights (ie that are the closest)
 *     - with the smallest weights
 *     - and with the smallest indices (ranks, numbered from 0) in `strng`
 *
 * Output:
 *   an array of two arrays, each subarray in the format:
 *     [number-weight, index in strng of the corresponding number, original number]
 *   The two subarrays are sorted in ascending order by their number weights if
 *   these weights are different, by their indexes in the string if equal.
 *
 * If n == 0, closest("") should return [].
 *
 * Examples:
 *   closest("103 123 4444 99 2000")    -> [[2, 4, 2000], [4, 0, 103]]
 *   closest("80 71 62 53")             -> [[8, 0, 80], [8, 1, 71]]
 *   closest("444 2000 445 544")        -> [[13, 2, 445], [13, 3, 544]]
 */

function closest(strng) {
  // your code here
}