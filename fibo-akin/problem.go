package main

/*
Fibo akin - 5 kyu

Be u(n) a sequence beginning with:
u[1] = 1, u[2] = 1, u[3] = 2, u[4] = 3, u[5] = 3, u[6] = 4,
u[7] = 5, u[8] = 5, u[9] = 6, u[10] = 6, u[11] = 6, u[12] = 8,
u[13] = 8, u[14] = 8, u[15] = 10, u[16] = 9, u[17] = 10, u[18] = 11,
u[19] = 11, u[20] = 12, u[21] = 12, u[22] = 12, u[23] = 12

The recurrence relation is:
u[n] = u[n - u[n-1]] + u[n - u[n-2]]

Task 1: Given two numbers n, k (integers > 2), write the function lengthSupUK(n, k)
returning the number of terms u[i] >= k with 1 <= i <= n.

Task 2: Given n (integer > 2), write the function comp(n)
returning the number of times where a term of u is less than its predecessor up to and including u[n].

Examples:
lengthSupUK(23, 12) => 4
lengthSupUK(50, 10) => 35
lengthSupUK(500, 100) => 304

comp(23) => 1
comp(100) => 22
comp(200) => 63
*/

// lengthSupUK returns the count of terms u[i] >= k for 1 <= i <= n
func lengthSupUK(n, k int) int {
	return 0
}

// comp returns the count of times u[i] < u[i-1] for 2 <= i <= n
func comp(n int) int {
	return 0
}