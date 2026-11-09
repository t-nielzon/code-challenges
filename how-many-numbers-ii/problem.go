package kata

/*
We want to find the numbers higher or equal than 1000 that the sum of every four consecutives digits cannot be higher than a certain given value.
If the number is num = d1d2d3d4d5d6, and the maximum sum of 4 contiguous digits is maxSum, then:
  d1 + d2 + d3 + d4 <= maxSum
  d2 + d3 + d4 + d5 <= maxSum
  d3 + d4 + d5 + d6 <= maxSum

We need to create a function, MaxSumDig(), that receives nMax, as the max value of the interval to study (the range [1000, nMax]), and a certain value, maxSum, the maximum sum that every four consecutive digits should be less or equal to.

The function should output a list with:
(1) the amount of numbers that satisfy the constraint
(2) the closest number to the mean of the results; if there are more than one, the smallest number should be chosen
(3) the total sum of all the found numbers

Example:
  MaxSumDig(2000, 3) returns [11, 1110, 12555]
  MaxSumDig(2000, 4) returns [21, 1120, 23665]
*/

func MaxSumDig(nMax, maxSum int) []int {
	// TODO: Implement solution
	return []int{}
}