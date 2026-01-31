/*
We search non-negative integer numbers, with at most 3 digits, such as the sum
of the cubes of their digits is the number itself; we will call them "cubic" numbers.

153 is such a "cubic" number : 1^3 + 5^3 + 3^3 = 153

These "cubic" numbers of at most 3 digits are easy to find, even by hand, so they
are "hidden" with other numbers and characters in a string.

The task is to find, or not, the "cubic" numbers in the string and then to make
the sum of these "cubic" numbers found in the string, if any, and to return a
string such as:

"number1 number2 (and so on if necessary) sumOfCubicNumbers Lucky"

if "cubic" numbers number1, number2, ... were found.

The numbers in the output are to be in the order in which they are encountered
in the input string.

If no cubic numbers are found return the string: "Unlucky".

Notes:
- In the string "001234" where 3 digits or more follow each other the first group
  to examine is "001" and the following is "234".
- When a continuous string of digits exceeds 3, the string is split into groups
  of 3 from the left. The last grouping could have 3, 2 or 1 digits.
*/
package kata

func IsSumOfCubes(s string) string {
}