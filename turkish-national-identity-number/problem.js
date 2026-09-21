/*
 * Turkish National Identity Number Validator
 *
 * Every Turkish citizen has an identity number whose validity can be checked by these set of rules:
 *
 * - It has 11 digits and the first digit can't be zero
 * - Take the sum of 1st, 3rd, 5th, 7th and 9th digit and multiply it by 7.
 *   Then subtract the sum of 2nd, 4th, 6th and 8th digits from this value.
 *   Modulo 10 of the result should be equal to the 10th digit
 * - Sum of first ten digits' modulo 10 should be equal to the 11th digit
 *
 * Task: Write a function to check if the given number is a valid Turkish ID number,
 * and return `true` or `false` accordingly. Return `false` for any invalid input too (e.g. strings).
 */

function isValidTurkishId(input) {

}