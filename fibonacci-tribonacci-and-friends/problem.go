/*
Fibonacci, Tribonacci and friends

Xbonacci is a generalization of the Fibonacci and Tribonacci sequences.
Given a signature (starting array of X elements), each following element
is the sum of the previous X elements.

Examples:
  xbonacci {1,1,1,1} 10 = {1,1,1,1,4,7,13,25,49,94}
  xbonacci {0,0,0,0,1} 10 = {0,0,0,0,1,1,2,4,8,16}
  xbonacci {1,0,0,0,0,0,1} 10 = {1,0,0,0,0,0,1,2,3,6}
  xbonacci {1,1} produces the Fibonacci sequence

Input:
  - signature: array of initial elements (can be any length)
  - n: number of elements to return

Output:
  - array containing the first n elements of the xbonacci sequence
*/

package main

func Xbonacci(signature []int, n int) []int {
	// TODO: implement
	return []int{}
}