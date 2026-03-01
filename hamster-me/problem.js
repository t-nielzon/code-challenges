/*
 * Hamster me (5 kyu)
 *
 * Write a function that accepts two inputs: code and message
 * and returns an encrypted string from message using the code.
 *
 * The code is a string that generates the key:
 * - All letters from code get number 1
 * - All letters directly following code letters get number 2
 *   (unless already assigned a smaller number), etc.
 * - Letters not reachable from any code letter wrap after z
 *   and continue from the last code letter alphabetically.
 *
 * Each letter encodes as: (root code letter)(distance number)
 *
 * Examples:
 *   hamsterMe('hamster', 'hamster') => 'h1a1m1s1t1e1r1'
 *   hamsterMe('hamster', 'helpme')  => 'h1e1h5m4m1e1'
 *
 * Notes:
 * - code will have at least 1 letter
 * - duplicates in code are possible and should be handled
 * - code and message consist of only lowercase letters
 */

function hamsterMe(code, message) {
}