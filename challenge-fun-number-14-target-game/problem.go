// ## Task
// In your favorite game, you must shoot a target with a water-gun to gain
// points. Each target can be worth a different amount of points.
//
// You are guaranteed to hit every target that you try to hit. You cannot hit
// consecutive targets though because targets are only visible for one second
// (one at a time) and it takes you a full second to reload your water-gun after
// shooting (you start the game already loaded).
//
// Given an array `vals` with the order of each target's point value, determine
// the maximum number of points that you can win.
//
// ## Example
// For `vals = [1, 2, 3, 4]`, the result should be `6`.
// (shoot vals[1]=2 and vals[3]=4)
//
// For `vals = [5, 5, 5, 5, 5]`, the result should be `15`.
// (shoot the 1st, 3rd and 5th)
//
// Note that the value can be zero or negative, don't shoot them.
//
// For `vals = [0, 0, -1, -1]`, the result should be `0`.
// For `vals = [5, -2, -9, -4]`, the result should be `5`.
//
// ## Input/Output
// - [input] integer array `vals`
//   The point values (negative or non-negative) of the targets (in order).
// - [output] an integer
//   The maximum number of points that you can score.

package kata

func TargetGame(vals []int) int {
	// your code here
}