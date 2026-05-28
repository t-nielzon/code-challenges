/*
 * Range of Integers in an Unsorted String
 * Difficulty: 5 kyu
 *
 * In this kata, your task is to write a function that returns the smallest and
 * largest integers in an unsorted string. A range is considered a finite
 * sequence of consecutive integers.
 *
 * Input:
 *  - A string comprised of integers in an unknown range; the result when a
 *    range of integers is shuffled around in random order then joined together.
 *  - An integer value representing the size of the range.
 *
 * Output:
 *  - Return the starting (minimum) and ending (maximum) numbers of the range
 *    as a slice of two integers.
 *
 * Example:
 *  MysteryRange("1568141291110137", 10) // [6 15]
 *
 * Technical Details:
 *  - The maximum size of a range will be 100 integers.
 *  - The starting number of a range will be: 0 < n < 100.
 */

package kata

func MysteryRange(s string, n int) [2]int {
}