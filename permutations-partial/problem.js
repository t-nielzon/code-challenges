/**
 * # Permutations (partial)
 *
 * Making a partial permutation is selecting items from a collection, where the order of the items does matter.
 * (This kata assumes all items in the collection are different.)
 *
 * Making a full permutation is selecting all items from a collection. Making a partial permutation is selecting
 * a number, between zero and the size of the collection, both inclusive, of items from a collection.
 *
 * Task: Implement an `Array` method `.permutations(n)` that returns a (possibly empty) `array` of all possible
 * unique (full) permutations of a combination of `n` items from an `array`.
 *
 * Requirements:
 * - Should throw an `Error` when an invalid number `n` is specified
 * - Inputs should not be modified
 * - The output should be a new `array`
 * - Valid inputs for `n` are non-negative integers
 *
 * Examples:
 * [ 0, 1 ].permutations(0) => [ [ ] ]
 * [ 0, 1 ].permutations(1) => [ [ 0 ], [ 1 ] ]
 * [ 0, 1 ].permutations(2) => [ [ 0, 1 ], [ 1, 0 ] ]
 * [ 0, 1 ].permutations(3) => [ ]
 */

Array.prototype.permutations = function(n) {
  // TODO: implement
};