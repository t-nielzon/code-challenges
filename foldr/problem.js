// # foldr - Introducing reduceRight with lazy evaluation
//
// JavaScript `reduceRight` will, given a `Function` and an initial value,
// reduce an `Array` to a single value by iterating it right to left and
// folding it with the given `Function`.
//
// The task is to define `Array` method `foldr`, which takes a function and
// an initial value as arguments, and returns a lazily evaluated fold over
// its `this`-argument.
//
// Also define `String` method `foldr`, which treats `String`s as `Array`s
// of characters.
//
// The key difference from eager evaluation is that the second argument to
// the folding function should be lazy - evaluated only if and when required.
// This allows the recursion to terminate early if the function doesn't
// evaluate its second argument.

Array.prototype.foldr = function(fn, z) {
  // implementation
};

String.prototype.foldr = function(fn, z) {
  // implementation
};