// Sum of a Sequence [Hard-Core Version]
// 
// The task is to sum all numbers from the beginning to the limit (possibly included),
// going in steps expressed by the third parameter.
//
// Rules:
// - If beginning equals limit, return it regardless of step
// - If impossible (beginning > limit with positive step, or beginning < limit with negative step), return 0
// - Step will never be 0
// - Very large sequences must use an efficient algorithm (no brute-force)
//
// Examples:
//   sumOfSequence(2, 2, 2)   => 2     (2)
//   sumOfSequence(2, 6, 2)   => 12    (2 + 4 + 6)
//   sumOfSequence(1, 5, 1)   => 15    (1 + 2 + 3 + 4 + 5)
//   sumOfSequence(1, 5, 3)   => 5     (1 + 4)
//   sumOfSequence(1, 5, -3)  => 0     (impossible)
//   sumOfSequence(-1, -5, 1) => 0     (impossible)
//   sumOfSequence(2, 2, -9)  => 2     (begin equals end)

function sumOfSequence(begin, end, step) {
  
}