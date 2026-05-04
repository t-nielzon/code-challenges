// Given a rational number n (n >= 0, denominator strictly positive) as a string
// (example: "2/3"), decompose this number as a sum of rationals with numerators
// equal to one and without repetitions (2/3 = 1/2 + 1/6 is correct but not
// 2/3 = 1/3 + 1/3, 1/3 is repeated).
//
// The algorithm must be "greedy", so at each stage the new rational obtained
// in the decomposition must have a denominator as small as possible. In this
// manner the sum of a few fractions in the decomposition gives a rather good
// approximation of the rational to decompose.
//
// Example:
//   decompose("21/23") should return ["1/2", "1/3", "1/13", "1/359", "1/644046"]
//
// Notes:
// 1) The rational given to decompose could be greater than one or equal to one,
//    in which case the first "fraction" will be an integer (with an implicit
//    denominator of 1).
// 2) If the numerator parses to zero, decompose returns [].
// 3) The number could also be a decimal which can be expressed as a rational.
//
// Ref: http://en.wikipedia.org/wiki/Egyptian_fraction

function decompose(n) {
  // your code here
}