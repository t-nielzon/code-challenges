package main

/*
Kata: The Supermarket Queue
Difficulty: 6 kyu

There is a queue for the self-checkout tills at the supermarket. Your task is write a function to calculate the total time required for all the customers to check out!

Input:
- customers: an array of positive integers representing the queue. Each integer represents a customer, and its value is the amount of time they require to check out.
- n: a positive integer, the number of checkout tills.

Output:
The function should return an integer, the total time required.

Examples:
queueTime([5,3,4], 1) => 12
queueTime([10,2,3,3], 2) => 10
queueTime([2,3,10], 2) => 12

Clarifications:
- There is only ONE queue serving many tills
- The order of the queue NEVER changes
- The front person in the queue proceeds to a till as soon as it becomes free
*/

func QueueTime(customers []int, n int) int {
}