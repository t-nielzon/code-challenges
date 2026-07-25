/*
The `fix` function can ( with lazy evaluation and a recursive `let` ) be defined as:

let fix = f => f (fix (f))

If we regard `fix` as a language primitive, any recursive function can be written without using recursion. For a more detailed explanation of `fix`, see Wikipedia on Fixed-point combinator.

Why?
* Some languages simply do not support direct recursion - a function is only available after it is defined. With `fix`, recursion can be attained without language support.
* By composing memoisation with the function and then `fix`ing it, a memoised recursive function may be obtained, without any demands on the base function.

Task:
* Refactor `fix` for use with strict evaluation.
* Refactor `factorial` and `fibonacci` as non-recursive functions that can be "fixed" to give recursive ones.
* Refactor `foldr` as a non-recursive, `fix`able function.

Specifications:
* `fix` accepts a non-recursive function and returns its recursive equivalent.
* `factorial` accepts a `BigInt` and returns its factorial as a `BigInt`.
* `fibonacci` accepts a non-negative `Number` and returns the corresponding Fibonacci number as a `BigInt`.
* `foldr` accepts a folding function, an initial value, and a list (Generator Object), and returns the list folded to a single value, right-associatively.
  - folding function: element -> thunk -> accumulator, where the thunk returns the accumulated result
  - initial value: a valid start value for the accumulator
  - list: a single-use Generator Object

All functions in this kata are curried.
*/

function fix(f) {
  // Implementation here
}

function factorial(f) {
  // Implementation here
}

function fibonacci(f) {
  // Implementation here
}

function foldr(f) {
  // Implementation here
}