package main

/*
We want to find the numbers higher or equal than 1000 that the sum of every four
consecutive digits cannot be higher than a certain given value.

If the number is num = d1d2d3d4d5d6, and the maximum sum of 4 contiguous digits is maxSum:
- d1 + d2 + d3 + d4 <= maxSum
- d2 + d3 + d4 + d5 <= maxSum
- d3 + d4 + d5 + d6 <= maxSum

Create a function MaxSumDig() that receives nMax (max value in range [1000, nMax]) and
maxSum (maximum sum of 4 consecutive digits).

The function returns a slice with:
- (1) count of numbers satisfying the constraint
- (2) the number closest to the mean (if tie, pick smallest)
- (3) the sum of all found numbers
*/

func MaxSumDig(nMax int, maxSum int) []int {
    return []int{}
}