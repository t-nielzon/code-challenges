/*
# foldr - Lazy Fold Right

Implement Array.prototype.foldr and String.prototype.foldr with lazy evaluation.

foldr takes a function and initial value, and folds them right to left.
The key feature is lazy evaluation: the second argument to fn is only evaluated
if and when required. This allows early termination of recursion before processing
the entire list.

Example:
  [1,2,3].foldr((a, b) => a + b(), 0) 
  // Recursively folds from left, passing accumulated result as thunk

String.prototype.foldr treats the string as an array of characters.
*/

Array.prototype.foldr = function(fn, z) {
  // TODO: implement lazy fold right
};

String.prototype.foldr = function(fn, z) {
  // TODO: implement lazy fold right for strings
};