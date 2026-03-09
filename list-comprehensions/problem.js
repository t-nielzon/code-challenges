/*
 * List Comprehensions (5 kyu)
 *
 * List comprehensions are a great feature provided by python.
 *
 * The purpose of this kata is to parse basic list comprehensions for javascript.
 *
 * Examples:
 *   var a = [1,2,3,4,5];
 *   [x for x in a]       // => [1,2,3,4,5]
 *   [x*x+2 for x in a]   // => [3,6,11,18,27]
 *   [17 for x in a]       // => [17,17,17,17,17]
 *
 * Write a function listComprehension(str) that evaluates a given list
 * comprehension in string form. Throw an error if the list comprehension
 * is invalid.
 *
 * Example:
 *   listComprehension("[3 for x in [3,1,2,3]]") // => [3,3,3,3]
 *
 * Note: the array part can be a function call that returns an array.
 */

function listComprehension(str) {
}

module.exports = listComprehension;