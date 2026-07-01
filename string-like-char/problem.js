/*
 * String like [Char]
 *
 * In some programming languages, strings internally are implemented like an
 * array of chars, which allows applying array methods to strings.
 *
 * Sadly, in JavaScript there is no Char type, and strings do not have most of
 * the methods of arrays.
 *
 * The objective of this kata is to add to the String.prototype the following
 * Array.prototype methods:
 *
 *   map, join, filter, forEach, some, every, reduce, reduceRight, sort, reverse
 *
 * Only methods that satisfy these two conditions were included:
 *   - There is not already a method for the String class with that name.
 *   - Since strings are immutable, the method must not modify the array
 *     (sort and reverse are exceptions: they mutate but also return the value).
 *
 * We also implement the mutable methods below, but instead of returning what
 * the array equivalent returns, they return the resulting string:
 *
 *   push, pop, shift, unshift, splice
 *
 * Examples:
 *   "Hello".push(" Peter"); // "Hello Peter"
 *   "Hello".pop();          // "Hell"
 */

// implementation is done by extending String.prototype; no single stub function.
// String.prototype.map = function (callback, thisArg) { };