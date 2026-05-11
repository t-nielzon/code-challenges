/*
 * Fix the pipes (5 kyu)
 *
 * Replace every `x` in the rectangular `map` with one of the heavy
 * box-drawing pipes (┗ ┓ ┏ ┛ ━ ┃) so that water entering on the
 * left side at row `start` flows out on the right side at row `end`.
 *
 * The pipeline must move only left-to-right (it never doubles back),
 * but it may travel up and down between columns. Cells that are not
 * `x` (e.g. `.`) must be left untouched.
 *
 * Unicode codepoints used for the pipes:
 *   ┗ - 9495 BOX DRAWINGS HEAVY UP AND RIGHT
 *   ┓ - 9491 BOX DRAWINGS HEAVY DOWN AND LEFT
 *   ┏ - 9487 BOX DRAWINGS HEAVY DOWN AND RIGHT
 *   ┛ - 9499 BOX DRAWINGS HEAVY UP AND LEFT
 *   ━ - 9473 BOX DRAWINGS HEAVY HORIZONTAL
 *   ┃ - 9475 BOX DRAWINGS HEAVY VERTICAL
 */

function fixThePipes(map, start, end) {
  // your code here
}