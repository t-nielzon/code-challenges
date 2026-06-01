/*
 * up AND down  (6 kyu)
 *
 * You are given a string `s` made up of substrings s(1), s(2), ..., s(n)
 * separated by whitespaces.
 * Example: "after be arrived two My so"
 *
 * Task
 * ----
 * Return a string `t` having the property (P):
 *
 *   length t(0) <= length t(1) >= length t(2) <= length t(3) >= length t(4) ...
 *
 * You must respect the following process rule:
 * at each step, from left to right, you can only move either already
 * consecutive strings or strings that became consecutive after a previous
 * move. The number of moves should be minimum. This process gives a UNIQUE
 * result (property P alone does not guarantee unicity).
 *
 * Walkthrough of the example:
 *   "after be arrived two My so"
 *   "after" > "be"        -> move  -> "be after arrived two My so"
 *   "after" < "arrived"   -> move  -> "be arrived after two My so"
 *   "after" > "two"       -> move  -> "be arrived two after My so"
 *   "after" > "My"        -> ok
 *   "My" == "so"          -> ok
 *
 * Once property (P) holds, to make the "up and down" visible:
 *   t(0), t(2), ...  -> lower case
 *   t(1), t(3), ...  -> upper case
 *
 * arrange("after be arrived two My so")  ->  "be ARRIVED two AFTER my SO"
 */

function arrange(s) {
  // your code here
}