/*
 * # Number encrypting: decypher
 * ## Part II of Number encrypting Katas
 *
 * You will receive a string cyphered according to the Kata "Number encrypting: cypher".
 * You must return the decyphered version of that string.
 *
 * Cypher table (number -> upper case / lower case letter):
 *   1: I / l
 *   2: R / z
 *   3: E / e
 *   4: A / a
 *   5: S / s
 *   6: G / b
 *   7: T / t
 *   8: B / (none)
 *   9: (none) / g
 *   0: O / o
 *
 * Any character not in the table does not change.
 *
 * Since one number can map to two letters, use these rules to pick the case:
 *   1. The first letter of the decyphered string is always Upper Case.
 *   2. Any letter after a period (period, then a space, then the letter) is Upper Case.
 *   3. Any letter 'i' (number 1) preceded by a space and followed by a space, a non-letter
 *      character, or the end of the string is Upper Case.
 *   4. In any other case, the converted letter is lower case. Existing upper case letters
 *      remain unchanged.
 *
 * Examples:
 *   "H3110 W0r1d"                                       -> "Hello World"
 *   "1 4m y0ur f47h3r"                                  -> "I am your father"
 *   "1 d0 n07 kn0w wh47 3153 1 c4n 7357. 83 c001. 600d 1uck" -> "I do not know what else I can test. Be cool. Good luck"
 *   "N3i7h3r d0 1"                                      -> "Neither do I"
 *   "1, 1. 1 m34n 13773r 1"                             -> "I, I. I mean letter I"
 *   "60475 4nd 60475"                                   -> "Goats and Boats"
 */

function decypher(str) {

}