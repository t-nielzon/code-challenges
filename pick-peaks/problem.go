/*
Pick peaks

In this kata, you will write a function that returns the positions and the values of the "peaks" (or local maxima) of a numeric array.

For example, the array arr = [0, 1, 2, 5, 1, 0] has a peak at position 3 with a value of 5 (since arr[3] equals 5).

The output will be returned as a PosPeaks struct with two fields: Pos and Peaks. Both of these fields should be slices. If there is no peak in the given array, then the output should be PosPeaks{Pos: []int{}, Peaks: []int{}}.

Example: PickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}) should return PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}}

All input arrays will be valid integer arrays (although it could still be empty), so you won't need to validate the input.

The first and last elements of the array will not be considered as peaks (in the context of a mathematical function, we don't know what is after and before and therefore, we don't know if it is a peak or not).

Also, beware of plateaus !!! [1, 2, 2, 2, 1] has a peak while [1, 2, 2, 2, 3] and [1, 2, 2, 2, 2] do not. In case of a plateau-peak, please only return the position and value of the beginning of the plateau. For example:
PickPeaks([]int{1, 2, 2, 2, 1}) returns PosPeaks{Pos: []int{1}, Peaks: []int{2}}
*/
package kata

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
}