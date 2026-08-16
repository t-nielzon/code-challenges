/*
 * The regular paperfolding sequence, also known as the dragon curve sequence,
 * is an infinite automatic sequence of 0s and 1s defined by the morphism:
 * 
 *   1 1 → 1 1 0 1
 *   0 1 → 1 0 0 1
 *   1 0 → 1 1 0 0
 *   0 0 → 1 0 0 0
 * 
 * Starting from [1, 1], each iteration applies the morphism to consecutive
 * pairs of bits, expanding the sequence progressively.
 * 
 * Task: Given an array of 0s and 1s, return if it is a leading substring
 * (prefix) of this sequence.
 */

function isPaperfoldSequence(arr) {
  // Your solution here
}