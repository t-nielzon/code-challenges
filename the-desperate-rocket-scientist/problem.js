/*
 * The Desperate Rocket Scientist (6 kyu)
 *
 * Jim's countdown function has been burned into ROM and cannot be changed:
 *
 *   function countdown() {
 *     var ret = "";
 *     var numbers = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
 *     for (var number in numbers) {
 *       if (ret) ret += ", ";
 *       if (number < 10)
 *         number = 10 - number;
 *       else if (number == 10)
 *         number = "Zero";
 *       ret += number;
 *     }
 *     ret += "!";
 *     return ret;
 *   }
 *
 * Jim's assistant Jeff introduced a bug elsewhere in the startup code
 * (most likely an enumerable property added to Array.prototype or
 * Object.prototype) that breaks the `for...in` loop above.
 *
 * Write a function `fix_countdown()` that, when called, repairs the
 * environment so that `countdown()` produces the expected output:
 *   "10, 9, 8, 7, 6, 5, 4, 3, 2, 1, Zero!"
 */

function fix_countdown() {
  // your code here
}