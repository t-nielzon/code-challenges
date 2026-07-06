/*
 * Foldr - Lazy Fold Right
 * 
 * Define Array method foldr, which takes a function and an initial value as arguments,
 * and returns a lazily evaluated fold over its this-argument.
 * Also define String method foldr, which treats Strings as Arrays of characters.
 * 
 * Unlike reduceRight which always processes the entire list, foldr with lazy evaluation
 * may end the recursion early if the folding function returns without evaluating its
 * second argument.
 */

Array.prototype.foldr = function(fn, z) {
  // TODO
};

String.prototype.foldr = function(fn, z) {
  // TODO
};