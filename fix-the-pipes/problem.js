/*
# Introduction
The goal of the kata is to connect water pipes from the water source to end of the pipe line without leaking anywhere.

# Task
Create a function which replaces all `x` in the `map` array with one of available pipes:
┗, ┓, ┏, ┛, ━, ┃ and returns new map as an output.
The water source is located on the left side of the map on `start` position (indexed from 0).
The end of pipe is located on the right side of the map on `end` position (indexed from 0).
The pipeline can go only from left side of the screen to the right side of the screen.
It never returns. It can go up and down.

# Notes
- The `map` creates always a rectangle.
- All inputs are correct. There is no need for any validation.

Unicode codepoints used for the pipes:
  ┗ - 9495 - BOX DRAWINGS HEAVY UP AND RIGHT
  ┓ - 9491 - BOX DRAWINGS HEAVY DOWN AND LEFT
  ┏ - 9487 - BOX DRAWINGS HEAVY DOWN AND RIGHT
  ┛ - 9499 - BOX DRAWINGS HEAVY UP AND LEFT
  ━ - 9473 - BOX DRAWINGS HEAVY HORIZONTAL
  ┃ - 9475 - BOX DRAWINGS HEAVY VERTICAL
*/

function fixThePipes(map, start, end) {
  // your code here
}