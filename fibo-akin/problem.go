/*
Fibo akin

Be u(n) a sequence beginning with:
u[1]  = 1,  u[2]  = 1,  u[3]  = 2,  u[4]  = 3,  u[5]  = 3,  u[6] = 4,
u[7]  = 5,  u[8]  = 5,  u[9]  = 6,  u[10] = 6,  u[11] = 6,  u[12] = 8,
u[13] = 8,  u[14] = 8,  u[15] = 10, u[16] = 9,  u[17] = 10, u[18] = 11,
u[19] = 11, u[20] = 12, u[21] = 12, u[22] = 12, u[23] = 12 etc...

The sequence is calculated as: u[n] = u[n - u[n-1]] + u[n - u[n-2]]

Task 1: Given n and k, return the count of terms u[i] >= k where 1 <= i <= n
Task 2: Given n, return the count of times u[i] < u[i-1] for 2 <= i <= n
*/

package main

// lengthSupUK returns the number of terms u[i] >= k with 1 <= i <= n
func lengthSupUK(n int, k int) int {
	return 0
}

// comp returns the number of times where u[i] < u[i-1] up to and including u[n]
func comp(n int) int {
	return 0
}