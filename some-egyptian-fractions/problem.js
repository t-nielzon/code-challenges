/*
 * Some Egyptian fractions
 *
 * Given a rational number n (n >= 0, with denominator strictly positive)
 * as a string (e.g. "2/3"), decompose this number as a sum of rationals
 * with numerators equal to one and without repetitions.
 *
 * The algorithm must be "greedy", so at each stage the new rational
 * obtained in the decomposition must have a denominator as small as possible.
 *
 * Notes:
 * - If the rational is >= 1, the first "fraction" will be an integer.
 * - If the numerator is zero, return [].
 * - The number could also be a decimal (e.g. "0.6").
 *
 * Example:
 *   decompose("21/23") => ["1/2", "1/3", "1/13", "1/359", "1/644046"]
 */
function decompose(nOverd) {
}

module.exports = { decompose };