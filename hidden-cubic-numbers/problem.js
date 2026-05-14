/*
 * Hidden "Cubic" numbers
 * Difficulty: 6 kyu
 *
 * We search non-negative integer numbers, with at most 3 digits, such as the sum
 * of the cubes of their digits is the number itself; we will call them "cubic" numbers.
 *
 *   153 is such a "cubic" number : 1^3 + 5^3 + 3^3 = 153
 *
 * These "cubic" numbers of at most 3 digits are easy to find, even by hand, so they
 * are "hidden" with other numbers and characters in a string.
 *
 * The task is to find, or not, the "cubic" numbers in the string and then to make
 * the sum of these "cubic" numbers found in the string, if any, and to return a
 * string such as:
 *   "number1 number2 (and so on if necessary) sumOfCubicNumbers Lucky"
 * if "cubic" numbers number1, number2, ... were found.
 *
 * The numbers in the output are to be in the order in which they are encountered
 * in the input string.
 *
 * If no cubic numbers are found return the string: "Unlucky".
 *
 * Examples:
 *   s = "aqdf&0#1xyz!22[153(777.777"
 *     groups: 0, 1, 22, 153, 777, 777
 *     cubic: 0, 1, 153 -> sum 154
 *     Return: "0 1 153 154 Lucky"
 *
 *   s = "QK29a45[&erui9026315"
 *     groups: 29, 45, 902, 631, 5 -> none cubic
 *     Return: "Unlucky"
 *
 * Notes:
 * - In "001234" the first group is "001" then "234".
 * - When a continuous run of digits exceeds 3, split from the left into groups of 3.
 */

function cubicNumbers(s) {
  // your code here
}