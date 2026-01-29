/*
 * String transformations
 * 
 * You are given a string and a transformation pattern (an array of operations).
 * Your task is to apply the operations in order to the string and return the final result.
 * 
 * Each operation in the pattern is one of:
 * 
 * "U" → convert the string to uppercase
 * "L" → convert the string to lowercase
 * "R" → reverse the string
 * "D" → duplicate each character (e.g., "abc" → "aabbcc", "qqbbcc" → "qqqqbbbbcccc")
 * 
 * Examples:
 * stringTransformation("abcd", ["U","R"]) // "DCBA"
 * stringTransformation("abc", ["D","L"])  // "aabbcc"
 * stringTransformation("AbCd", ["L","R"]) // "dcba"
 * 
 * Note:
 * Input string length is 1 <= n <= 10^5, length of transformations is 0 <= n <= 10^5,
 * in all tests is guaranteed that the string will be not longer than 10^8.
 */

function stringTransformation(s, transformations) {
  // Your solution here
}

module.exports = { stringTransformation };