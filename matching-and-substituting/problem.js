/*
 * Matching And Substituting
 *
 * I got lots of files beginning like this:
 *
 *   Program title: Primes
 *   Author: Kern
 *   Corporation: Gold
 *   Phone: +1-503-555-0091
 *   Date: Tues April 9, 2005
 *   Version: 6.7
 *   Level: Alpha
 *
 * Here we will work with strings like the string data above and not with files.
 *
 * The function change(s, prog, version) given:
 *   s=data, prog="Ladder", version="1.1"
 * will return:
 *   "Program: Ladder Author: g964 Phone: +1-503-555-0090 Date: 2019-01-01 Version: 1.1"
 *
 * Rules:
 * - The date should always be "2019-01-01".
 * - The author should always be "g964".
 * - Replace the current "Program Title" with the prog argument. Also remove
 *   "Title", so "Program Title: Primes" becomes "Program: Ladder".
 * - Remove the lines containing "Corporation" and "Level" completely.
 * - Phone numbers and versions must be in valid formats.
 *   A valid version is one or more digits, a dot, one or more digits
 *   (0.6, 5.4, 14.275, 1.99 valid; .6, 5, 14.2.7, 1.9.9 invalid).
 *   A valid phone is +1-xxx-xxx-xxxx where each x is a digit.
 * - If the phone or version format is not valid, return "ERROR: VERSION or PHONE".
 * - If the version is valid and is anything other than "2.0", replace it with
 *   the version parameter. If it's "2.0", don't modify it.
 * - If the phone is valid, replace it with "+1-503-555-0090".
 */

function change(s, prog, version) {
  // your code here
}