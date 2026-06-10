/*
 * # foldr
 *
 * ## Introducing reduceRight, version -1
 *
 * JavaScript `reduceRight` will, given a Function and an initial value,
 * reduce an Array to a single value by iterating it, and the initial
 * value, right to left and folding it with the given Function.
 *
 * Given that Arrays will always be finite and JavaScript evaluation is
 * strict, this is functionally equivalent to a recursive `foldr` where
 * the arguments of `fn` are reversed.
 *
 * ## Introducing laziness
 *
 * Now imagine that `fn` arguments are evaluated lazily - evaluated only
 * if and when required. Where `reduceRight` will always process the
 * entire list, `foldr` with lazy evaluation may arrive at the base case
 * and end the recursion before having processed the entire list.
 *
 * This kata's scope is ending the recursion when appropriate, possibly
 * before processing the entire list, by being lazy.
 *
 * ## Task
 *
 * Define Array method `foldr`, which takes a function and an initial
 * value as arguments, and returns a lazily evaluated fold over its
 * `this`-argument. The second argument passed to `fn` is a thunk that,
 * when called, lazily computes the fold of the rest of the list.
 *
 * Also define String method `foldr`, which treats Strings as Arrays of
 * characters.
 *
 * There will be tests with Arrays of Numbers, characters (Strings) and
 * Booleans.
 */

Array.prototype.foldr = function (fn, z) {
  // your code here
};

String.prototype.foldr = Array.prototype.foldr;