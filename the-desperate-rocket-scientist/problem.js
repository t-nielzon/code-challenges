/*
 * The Desperate Rocket Scientist
 *
 * Jim, the rocket scientist, has finished the code for the board computer of
 * his new Mars rocket. Only one last function is missing: the function for
 * creating the countdown.
 *
 * Today is rocket launch day, and the unit test for the countdown suddenly
 * fails. Jim swears he didn't change anything in the countdown function. He
 * suspects his assistant Jeff introduced a bug somewhere in the rocket board
 * computer startup code, but cannot understand how that could affect the
 * output of his countdown method in such a strange way.
 *
 * The countdown function has already been burned into a ROM that cannot be
 * replaced. Jim has to fix the problem by calling a method in the startup
 * code that is still accessible.
 *
 * Here is Jim's countdown code that he is not able to change any more:
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
 * Write a function `fix_countdown()` that contains a fix to make `countdown()`
 * work again.
 *
 * The bug: Jeff added enumerable properties to Array.prototype, so the
 * `for...in` loop in countdown() now iterates over those extra properties too.
 */

function fix_countdown() {
  // your code here
}