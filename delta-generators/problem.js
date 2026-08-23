/*
Delta Generators

In mathematics, the symbols Δ and d are often used to denote the difference between two values. 
Similarly, differentiation takes the ratio of changes (ie. dy/dx) for a linear relationship. 
This method can be applied multiple times to create multiple 'levels' of rates of change.
(A common example is x (position) -> v (velocity) -> a (acceleration)).

Our function `delta` will take a sequence of `values` and a positive integer `level`, 
and return a sequence with the 'differences' of the original `values`. 
(Differences here means strictly b - a, eg. [1, 3, 2] => [2, -1])

The `level` is the 'level' of difference. For example:
  input = [1, 2, 4, 7, 11, 16, 22]
  delta(input, 1) => [1, 2, 3, 4, 5, 6]
  delta(input, 2) => [1, 1, 1, 1, 1]
  delta(input, 3) => [0, 0, 0, 0]

Input and output can be any iterable, possibly infinite. The function must work with:
- Both finite and infinite iterables
- Types that support the subtraction operator (number, BigInt)
- Generator objects and arrays

Each subsequent level will be one item shorter than the previous.
If an infinite input is provided, the output must also be infinite.
*/

function* delta(values, level) {
  // Your solution here
}