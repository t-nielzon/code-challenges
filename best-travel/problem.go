package main

/*
John and Mary want to travel between a few towns. Mary has a list of distances
between towns. John doesn't want to drive more than t miles and will visit only
k towns. The task is to find the maximum sum of k distances that doesn't exceed t.

Example:
ls = [50, 55, 57, 58, 60], t = 174, k = 3
Possible sums: 162, 163, 165, 165, 167, 168, 170, 172, 173, 175
Maximum sum ≤ 174 is 173 (distances [55, 58, 60])

Parameters:
- t: maximum sum of distances (integer >= 0)
- k: number of towns to visit (k >= 1)
- ls: list of distances (all non-negative integers)

Returns:
- The largest sum of k distances ≤ t, or -1 if no such sum exists
*/
func chooseBestSum(t int, k int, ls []int) int {
}