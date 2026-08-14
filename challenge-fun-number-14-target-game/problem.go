package main

/*
Challenge Fun #14: Target Game

In your favorite game, you must shoot a target with a water-gun to gain points.
Each target can be worth a different amount of points.

You are guaranteed to hit every target that you try to hit. You cannot hit
consecutive targets though because targets are only visible for one second
(one at a time) and it takes you a full second to reload your water-gun after
shooting (you start the game already loaded).

Given an array vals with the order of each target's point value, determine
the maximum number of points that you can win.

Examples:
- vals = [1, 2, 3, 4] -> 6 (shoot indices 1 and 3: 2 + 4)
- vals = [5, 5, 5, 5, 5] -> 15 (shoot indices 0, 2, 4: 5 + 5 + 5)
- vals = [0, 0, -1, -1] -> 0 (don't shoot anything)
- vals = [5, -2, -9, -4] -> 5 (shoot index 0 only)
*/

func maxScore(vals []int) int {
}