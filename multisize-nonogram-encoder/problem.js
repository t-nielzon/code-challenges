/*
Multisize Nonogram Encoder
Difficulty: 5 kyu

If you're not familiar with nonograms, check: https://en.wikipedia.org/wiki/Nonogram

Given a solved nonogram (represented as an array of arrays with 0s and 1s),
convert it into a tuple of (column_clues, row_clues).

Each clue is an array of consecutive 1s in that row or column.
For example:
  - [0, 0, 1, 1, 1, 0, 1, 1, 1, 1] -> [3, 4]
  - [1, 1, 1, 0, 0, 1, 1] -> [3, 2]

Input: nonogram - a 2D array of 0s and 1s (5 <= size <= 100)
Output: [columnClues, rowClues] where each clue is an array of integers

Notes:
- Empty rows/columns should be represented as empty arrays []
- There are 200 random tests
*/

function nonogramEncoder(nonogram) {
  
}