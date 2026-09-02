package main

// Connecting Values
// Given a two-dimensional array of non-negative integers arr, a value val, and a coordinate coord in the form (row, column),
// return an iterable of all of the coordinates that contain the given value and are connected to the original coordinate
// by the given value. A connection to the original coordinate is considered valid as long as there is a path of pairwise
// connected cells, horizontally, vertically, or diagonally.
//
// If the value of arr at coord is not equal to val, return an empty iterable.
// The coordinates must include the original coordinate and may be in any order.
//
// Examples:
// Array: [1,0,2,0,2,1]
//        [1,0,2,1,5,7]
//        [4,1,1,0,1,9]
//
// With val 1 and coord (0, 0), output: [(2, 4), (2, 1), (0, 0), (2, 2), (1, 0), (1, 3)]
// With val 2 and coord (0, 2), output: [(0, 2), (1, 2)]
// With val 0 and coord (0, 0), output: [] (empty)

func ConnectingValues(arr [][]int, val int, coord []int) [][]int {
	return [][]int{}
}