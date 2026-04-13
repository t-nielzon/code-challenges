package kata

func buildSequence(n int) []int {
	u := make([]int, n+1)
	if n >= 1 {
		u[1] = 1
	}
	if n >= 2 {
		u[2] = 1
	}
	for i := 3; i <= n; i++ {
		u[i] = u[i-u[i-1]] + u[i-u[i-2]]
	}
	return u
}

func LengthSupUK(n, k int) int {
	u := buildSequence(n)
	count := 0
	for i := 1; i <= n; i++ {
		if u[i] >= k {
			count++
		}
	}
	return count
}

func Comp(n int) int {
	u := buildSequence(n)
	count := 0
	for i := 2; i <= n; i++ {
		if u[i] < u[i-1] {
			count++
		}
	}
	return count
}