/*
 * Ecco the Dolphin - Ecco Junior
 *
 * Guide Ecco the dolphin through a 2D underwater maze to reach Tara.
 * The maze is a string with rows separated by '\n'.
 *
 * Input:
 *   A string representing the underwater maze.
 *
 * Output:
 *   An array of [y, x] tuples representing the path from Ecco to Tara.
 *   (0, 0) is the top-left corner; y goes down, x goes right.
 *   If no path is possible, return null.
 *
 * Characters:
 *   ~  Sea (swimmable)
 *   .  Air (unswimmable)
 *   #  Rock (obstacle)
 *   e  Ecco (start)
 *   t  Tara (goal)
 *   S  Shark   - moves horizontally, bouncing off rocks
 *   Y  Jellyfish- moves vertically, bouncing off rocks and air
 *   O  Urchin  - expands to all 8 adjacent tiles (diagonals included)
 *
 * Rules:
 *   1. Ecco swims in the 4 cardinal directions only (no diagonals).
 *   2. Ecco only swims through sea (~) tiles.
 *   3. Ecco cannot pass rocks (#), air (.), or leave the grid.
 *   4. Ecco avoids any tile a creature might occupy:
 *        - Sharks occupy their horizontal lane, bounded by rocks.
 *        - Jellyfish occupy their vertical lane, bounded by rocks and air.
 *        - Urchins occupy all 8 adjacent tiles plus their own.
 *        - Creatures do not block each other.
 *   5. Ecco and Tara always start in a safe zone.
 *   6. All rows have equal length.
 *   7. Every creature (and Ecco/Tara) starts on a sea tile.
 *
 * Win: a valid path from e to t exists. Lose: no valid path -> null.
 */

function findPath(maze) {
  // your code here
}