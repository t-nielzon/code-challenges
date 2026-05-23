/*
 * All that is open must be closed...
 *
 * Background:
 * We all know about "balancing parentheses" (plus brackets, braces and chevrons)
 * and even balancing characters that are identical.
 *
 * Kata:
 * Write a function to validate that a supplied string is balanced.
 * You must determine if all that is open is then closed, and nothing is closed
 * which is not already open!
 *
 * You will be given a string to validate, and a second string, where each pair
 * of characters defines an opening and closing sequence that needs balancing.
 *
 * You may assume that the second string always has an even number of characters.
 *
 * Examples:
 *   ("(Sensei says yes!)", "()")     => true
 *   ("(Sensei says no!", "()")       => false
 *   ("(Sensei [says] yes!)", "()[]") => true
 *   ("(Sensei [says) no!]", "()[]")  => false
 *   ("Sensei says 'yes'!", "''")     => true
 *   ("Sensei say's no!", "''")       => false
 */

function isBalanced(s, caps) {
  // your code here
}