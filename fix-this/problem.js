/*
The `fix` function can ( with lazy evaluation and a recursive `let` ) be defined as:

let fix = f => f (fix (f))

If we regard `fix` as a language primitive, any recursive function can be written without using recursion. For a more detailed explanation of `fix`, see Wikipedia.

Unfortunately, your language has strict evaluation and no recursive `let` ( also called `let rec` ).

Why?
* Some languages simply do not support direct recursion - a function is only available after it is defined. With `fix`, recursion can be attained without language support.
* By composing memoisation with the function and then fixing it, a memoised recursive function may be obtained, without any demands on the base function. ( This will be shown in the tests with `fibonacci`. )

Task:
* Refactor `fix` for use with strict evaluation.
* Refactor `factorial` and `fibonacci` as non-recursive functions that can be "fixed" to give recursive ones. ( No optimisations necessary, just the trivial implementations ).
* Refactor `foldr` as a non-recursive, `fix`able function. This will need to support laziness in the second argument of the folding function. It will also need to accept a `Generator Object` instead of an `Array` to work on, so infinite lists can be encoded.

Specifications:
- `fix` accepts a non-recursive function and returns its recursive equivalent.
- `factorial` accepts a `BigInt` and returns its factorial as a `BigInt`.
- `fibonacci` accepts a non-negative `Number` and returns the corresponding ( `0,1-` ) Fibonacci number as a `BigInt`.
- `foldr` accepts a folding function, an initial value ( non-optional in this kata ), and a list, and returns the list folded to a single value, right-associatively.
  - folding function : element -> accumulator -> accumulator, where the second argument is a thunk for laziness
  - initial value : a valid start value for the accumulator
  - a list, encoded as a single-use Generator Object. The list may be infinite.

All functions are curried.

Type signatures:
- fix : (x -> x) -> x
- factorial : BigInt -> BigInt
- fibonacci : Number -> BigInt
- foldr : (x -> (() -> z) -> z) -> z -> [x] -> z
*/

const fix = f => {
  // TODO: implement
};

const factorial = f => n => {
  // TODO: implement
};

const fibonacci = f => n => {
  // TODO: implement
};

const foldr = self => folding_fn => init => gen => {
  // TODO: implement
};