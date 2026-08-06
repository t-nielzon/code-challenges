/*
 * ## groupBy
 *
 * Array elements are separated into (non-empty) groups of equal elements.
 * This equality can be according to the builtin `===` or according to a user-provided function,
 * which gives the option of grouping elements that are not strictly equal but are equal in some
 * attribute or property.
 *
 * Task: Implement an `Array` method `groupBy` that returns an `array` of non-empty `array`s
 * of elements that are equal according to the optional comparing `function`.
 *
 * The comparing function has the same signature as `===` : `(Value,Value) => Boolean`.
 * Inputs must not be modified.
 * The output is to be a new `array` of new `array`s.
 * The prototype method must not be enumerable.
 *
 * Examples:
 * [ 1,1,1 ].groupBy() => [ [1,1,1] ]
 * [ 0,1,1,0,0,1 ].groupBy() => [ [0], [1,1], [0,0], [1] ]
 *
 * function eqOdd(m,n) { return m%2===n%2; }
 * [ 1,3,2,2,4,1 ].groupBy(eqOdd) => [ [1,3], [2,2,4], [1] ]
 */

Object.defineProperty(Array.prototype, 'groupBy', {
  value: function(comparator) {
    
  },
  enumerable: false
});