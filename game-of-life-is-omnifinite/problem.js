/*
 * Game of Life is Omnifinite
 * Difficulty: 5 kyu
 *
 * Conway's Game of Life is a 2D cellular automaton. A life may die out after
 * any number of steps, making Game of Life "omnifinite".
 *
 * Objective:
 *   Given a non-negative integer N, return a life that fully dies out after
 *   exactly N steps, assuming an infinite grid. Represent it as a list of
 *   2D points [x, y]. Tested against N up to 2000.
 *
 * Notes:
 *   - After N steps every cell must be dead (no gliders, no still lifes).
 *   - Fewer total live cells across all generations is better (avoid timeouts).
 *
 * Examples:
 *   N=2 -> a diagonal segment of length 3 collapses to a single cell, then dies.
 *   N=3 -> a small symmetric pattern that fully clears after 3 generations.
 */

function omnifinite(n) {
  // your code here
}