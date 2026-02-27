/*
### Background

Pat Programmer is worried about the future of jobs in programming because of the advance of AI tools
like ChatGPT, and he decides to become a chef instead! So he wants to be an expert at flipping pancakes.

A pancake is characterized by its diameter, a positive integer.
Given a stack of pancakes, you can insert a spatula under any pancake and then flip, which reverses the
order of all the pancakes on top of the spatula.

### Task

Write a function that takes a stack of pancakes and returns a sequence of flips that arranges them in
order by diameter, with the largest pancake at the bottom and the smallest at the top. The pancakes are
given as a list of positive integers, with the left end of the list being the top of the stack.

A flip at position k reverses the subarray from index 0 to k (inclusive).

### Example

Consider the stack [1,5,8,3]. One way of achieving the desired order is as follows:
Flip at position 2 -> [8,5,1,3], flip at position 3 -> [3,1,5,8], flip at position 1 -> [1,3,5,8].
Return [2,3,1].

### Note

You don't have to find the shortest sequence of flips. However, don't include unnecessary flips:
1. No two consecutive flips that leave the stack in the same state.
2. Flipping only the top pancake (position 0) doesn't achieve anything.
*/
package kata

func Solve(stack []int) []int {
}