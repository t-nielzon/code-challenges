package main

func lengthSupUK(n int, k int) int {
	u := make([]int, n+1)
	u[1] = 1
	u[2] = 1
	
	count := 0
	if u[1] >= k {
		count++
	}
	if u[2] >= k {
		count++
	}
	
	for i := 3; i <= n; i++ {
		u[i] = u[i-u[i-1]] + u[i-u[i-2]]
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
	
	count := 0
	
	for i := 3; i <= n; i++ {
		u[i] = u[i-u[i-1]] + u[i-u[i-2]]
		if u[i] < u[i-1] {
			count++
		}
	}
	
	return count
}