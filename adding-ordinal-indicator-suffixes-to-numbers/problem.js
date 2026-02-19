/*
 * Adding ordinal indicator suffixes to numbers
 *
 * Complete the function which should take a number and return it as a string
 * with the correct ordinal indicator suffix (in English). That is:
 *
 *   1 ==> "1st"
 *   2 ==> "2nd"
 *   3 ==> "3rd"
 *   4 ==> "4th"
 *   ... and so on
 *
 * For the purposes of this kata, you may assume that the function will always
 * be passed a non-negative integer. If the function is given 0 as an argument,
 * it should return "0" (as a string).
 *
 * Rules:
 * - st is used with numbers ending in 1 (e.g. 1st)
 * - nd is used with numbers ending in 2 (e.g. 92nd)
 * - rd is used with numbers ending in 3 (e.g. 33rd)
 * - All "teen" numbers ending with 11, 12 or 13 use -th (e.g. 11th, 112th)
 * - th is used for all other numbers (e.g. 9th)
 */

function numberToOrdinal(n) {
}

module.exports = numberToOrdinal;