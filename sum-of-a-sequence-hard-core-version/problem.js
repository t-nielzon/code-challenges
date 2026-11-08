/*
 * Sum of a Sequence [Hard-Core Version]
 * Difficulty: 6 kyu
 * 
 * The task is to sum all numbers from `begin` to `end` (possibly included),
 * incrementing by `step`.
 * 
 * Rules:
 * - If begin equals end, return it regardless of step
 * - If sequence is impossible (e.g., begin > end with positive step), return 0
 * - step will never be 0
 * - Very large sequences will be tested; brute-force will not work
 * 
 * Examples:
 * begin=2, end=2, step=2   => 2
 * begin=2, end=6, step=2   => 12 (2 + 4 + 6)
 * begin=1, end=5, step=1   => 15 (1 + 2 + 3 + 4 + 5)
 * begin=1, end=5, step=3   => 5 (1 + 4)
 * begin=1, end=5, step=-3  => 0
 * begin=-1, end=-5, step=1 => 0
 * begin=2, end=2, step=-9  => 2
 */

function sumOfSequence(begin, end, step) {
  
}