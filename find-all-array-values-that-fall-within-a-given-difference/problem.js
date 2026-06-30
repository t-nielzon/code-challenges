/*
 * Find All Array Values That Fall Within a Given Difference
 *
 * You are given an array of non-negative integers.
 *
 * The goal is to find all the values in the array that are within a
 * given range (difference) of each other.
 *
 * Return the final values in ascending order.
 *
 * Example:
 *   numbers = [5, 32, 5, 1, 31, 70, 30, 8]
 *   difference = 2
 *   new GroupByDifference(numbers).find(2) // => [5, 5, 30, 31, 32]
 *
 *   new GroupByDifference([5, 32, 5, 1, 31, 70, 30, 8]).find(3)
 *     // => [5, 5, 8, 30, 31, 32]
 *
 * If an empty array is given, then an empty array should be returned
 * regardless of the difference value passed in.
 */

class GroupByDifference {
  constructor(numbers) {
    this.numbers = numbers;
  }

  find(difference) {
    // your code here
  }
}