package main

/*
In this kata, you will write a function that returns the positions and the values of the "peaks" (or local maxima) of a numeric array.

For example, the array [0, 1, 2, 5, 1, 0] has a peak at position 3 with a value of 5.

The output will be returned as an object with two properties: pos and peaks. Both should be arrays.
If there is no peak in the given array, return {pos: [], peaks: []}.

Important notes:
- The first and last elements are never considered peaks.
- Be aware of plateaus! [1, 2, 2, 2, 1] has a peak, but [1, 2, 2, 2, 3] does not.
- For plateau-peaks, return only the position and value of the beginning.

Example: pickPeaks([3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3]) returns {pos: [3, 7], peaks: [6, 3]}
*/

type Result struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(arr []int) Result {
	// implement here
}