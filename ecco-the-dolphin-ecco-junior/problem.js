/*
Ecco the Dolphin - Ecco Junior (5 kyu)

Guide Ecco (e) through a 2D underwater maze to reach Tara (t).

Input:
  A string representing the underwater maze (rows separated by '\n').

Output:
  An array of [y, x] tuples representing the path from Ecco to Tara.
  (0, 0) is top-left, y increases downward, x increases rightward.
  If no path is possible, return null.

Tiles:
  ~  Sea (swimmable)
  .  Air (unswimmable)
  #  Rock (obstacle)
  e  Ecco (start)
  t  Tara (goal)
  S  Shark - moves horizontally, bounces off rocks.
  Y  Jellyfish - moves vertically, bounces off rocks and air.
  O  Urchin - expands to all 8 adjacent tiles (including diagonals).

Rules:
  - Ecco swims in 4 cardinal directions only (no diagonals).
  - Ecco can only traverse sea tiles and cannot enter any tile/lane
    where a creature might potentially be.
  - Ecco and Tara always start in a safe zone.
  - All rows have equal length. Creature tiles count as sea.

Constraints:
  - Height <= 40, Width <= 100.
*/

function findPath(maze) {
  // your code here
}