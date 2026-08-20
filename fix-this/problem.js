/*
Fix this

The `fix` function can (with lazy evaluation and a recursive `let`) be defined as:
  let fix = f => f (fix (f))

If we regard `fix` as a language primitive, any recursive function can be written
without using recursion.

Unfortunately, JavaScript has strict evaluation and no recursive `let` (also called
`let rec`).

Task:
- Refactor `fix` for use with strict evaluation.
- Refactor `factorial` and `fibonacci` as non-recursive functions that can be "fixed"
  to give recursive ones.
- Refactor `foldr` as a non-recursive, `fix`able function. This will need to support
  laziness in the second argument of the folding function. It will also need to accept
  a Generator Object instead of an Array to work on.

All functions in this kata are curried.

Type signatures:
- fix : (x -> x) -> x
- factorial : BigInt -> BigInt
- fibonacci : Number -> BigInt
- foldr : (x -> (() -> z) -> z) -> z -> [x] -> z
*/

const fix = f => {
  // Your code here
};

const factorial = f => n => {
  // Your code here
};

const fibonacci = f => n => {
  // Your code here
};

const foldr = rec => f => acc => gen => {
  // Your code here
};