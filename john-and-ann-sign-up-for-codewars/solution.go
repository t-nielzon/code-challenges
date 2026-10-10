package solution

func ann(n int) []int {
	if n <= 0 {
		return []int{}
	}
	
	a := make([]int, n)
	j := make([]int, n)
	
	a[0] = 1
	j[0] = 0
	
	for i := 1; i < n; i++ {
		j[i] = i - a[j[i-1]]
		a[i] = i - j[a[i-1]]
	}
	
	return a
}

func john(n int) []int {
	if n <= 0 {
		return []int{}
	}
	
	a := make([]int, n)
	j := make([]int, n)
	
	a[0] = 1
	j[0] = 0
	
	for i := 1; i < n; i++ {
		j[i] = i - a[j[i-1]]
		a[i] = i - j[a[i-1]]
	}
	
	return j
}

func sum_ann(n int) int {
	if n <= 0 {
		return 0
	}
	
	a := make([]int, n)
	j := make([]int, n)
	
	a[0] = 1
	j[0] = 0
	sum := 1
	
	for i := 1; i < n; i++ {
		j[i] = i - a[j[i-1]]
		a[i] = i - j[a[i-1]]
		sum += a[i]
	}
	
	return sum
}

func sum_john(n int) int {
	if n <= 0 {
		return 0
	}
	
	a := make([]int, n)
	j := make([]int, n)
	
	a[0] = 1
	j[0] = 0
	sum := 0
	
	for i := 1; i < n; i++ {
		j[i] = i - a[j[i-1]]
		a[i] = i - j[a[i-1]]
		sum += j[i]
	}
	
	return sum
}