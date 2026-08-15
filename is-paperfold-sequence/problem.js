/**
 * The regular paperfolding sequence (dragon curve sequence) is defined by the morphism:
 *   1 1 → 1 1 0 1
 *   0 1 → 1 0 0 1
 *   1 0 → 1 1 0 0
 *   0 0 → 1 0 0 0
 * 
 * Starting with [1, 1], apply the morphism to consecutive pairs to generate:
 * [1, 1] → [1, 1, 0, 1] → [1, 1, 0, 1, 1, 0, 0, 1] → ...
 * 
 * Task: Given an array of 0s and 1s, return whether it is a prefix of this sequence.
 * 
 * Performance: O(n) time complexity required. Must handle arrays up to 100,000,000 elements.
 */
function isPaperfold(arr) {
  
}