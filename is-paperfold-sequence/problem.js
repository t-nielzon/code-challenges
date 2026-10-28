// problem.js
/*
 * The regular paperfolding sequence (also known as the dragon curve sequence)
 * is an infinite automatic sequence of 0s and 1s defined by the morphism:
 * 
 *   1 1 → 1 1 0 1
 *   0 1 → 1 0 0 1
 *   1 0 → 1 1 0 0
 *   0 0 → 1 0 0 0
 * 
 * Starting from [1, 1], repeatedly applying these rules generates:
 * [1, 1] → [1, 1, 0, 1] → [1, 1, 0, 1, 1, 0, 0, 1] → ...
 * 
 * Given an array of 0s and 1s, return if it is a leading substring (prefix)
 * of this sequence.
 * 
 * Performance constraint: O(n) time complexity required.
 */

function isPaperfoldSequence(arr) {
  
}