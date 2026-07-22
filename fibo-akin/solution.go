package main

func lengthSupUK(n, k int) int {
	u := make([]int, n+1)
	u[1] = 1
	u[2] = 1
	
	for i := 3; i <= n; i++ {
		u[i] = u[i-u[i-1]] + u[i-u[i-2]]
	}
	
	count := 0
	for i := 1; i <= n; i++ {
		if u[i] >= k {
			count++
		}
	}
	return count
}

func comp(n int) int {
	u := make([]int, n+1)
	u[1] = 1
	u[2] = 1
	
	for i := 3; i <= n; i++ {
		u[i] = u[i-u[i-1]] + u[i-u[i-2]]
	}
	
	count := 0
	for i := 2; i <= n; i++ {
		if u[i] < u[i-1] {
			count++
		}
	}
	return count
}