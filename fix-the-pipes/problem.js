/*
 * # Introduction
 *
 * The goal of the kata is to connect water pipes from the water source to
 * end of the pipe line without leaking anywhere.
 *
 * # Task
 *
 * Create a function which replaces all `x` in the `map` array with one of
 * available pipes : ┗ ┓ ┏ ┛ ━ ┃ and returns new map as an output.
 * The water source is located on the left side of the map on `start`
 * position (indexed from 0).
 * The end of pipe is located on the right side of the map on `end`
 * position (indexed from 0).
 * The pipeline can go only from left side of the screen to the right side
 * of the screen. It never returns. Of course it can also go up and down.
 *
 * This is not allowed (the pipeline goes back to the left):
 *   ━━━┓
 *   ┏━━┛
 *   ┗━━━
 *
 * # Correct samples
 *
 * Sample 1: start = 1, end = 1
 *   map:        solution:
 *   ...         ...
 *   xxx   -->   ━━━
 *   ...         ...
 *
 * Sample 2: start = 0, end = 0
 *   map:        solution:
 *   xxxx        ┓┏┓┏
 *   xxxx  -->   ┃┃┃┃
 *   xxxx        ┗┛┗┛
 *
 * Sample 3: start = 2, end = 2
 *   map:        solution:
 *   xxxx        ┏┓┏┓
 *   xxxx  -->   ┃┃┃┃
 *   xxxx        ┛┗┛┗
 *
 * # Notes
 * - The `map` creates always a rectangle.
 * - All inputs are correct. There is no need for any validation.
 *
 * Unicode codepoints for the pipes:
 *   ┗ 9495 | ┓ 9491 | ┏ 9487 | ┛ 9499 | ━ 9473 | ┃ 9475
 */

function fixThePipes(map, start, end) {
  // your code here
}