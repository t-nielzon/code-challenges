package main

func ann(n int) []int {
	if n == 0 {
		return []int{}
	}
	
	a := make([]int, n)
	j := make([]int, n)
	
	a[0] = 1
	j[0] = 0
	
	for day := 1; day < n; day++ {
		t_ann := a[day-1]
		t_john := j[day-1]
		
		a[day] = day - j[t_ann]
		j[day] = day - a[t_john]
	}
	
	return a
}

func john(n int) []int {
	if n == 0 {
		return []int{}
	}
	
	a := make([]int, n)
	j := make([]int, n)
	
	a[0] = 1
	j[0] = 0
	
	for day := 1; day < n; day++ {
		t_ann := a[day-1]
		t_john := j[day-1]
		
		a[day] = day - j[t_ann]
		j[day] = day - a[t_john]
	}
	
	return j
}

func sum_ann(n int) int {
	if n == 0 {
		return 0
	}
	
	a := make([]int, n)
	j := make([]int, n)
	
	a[0] = 1
	j[0] = 0
	sum := 1
	
	for day := 1; day < n; day++ {
		t_ann := a[day-1]
		t_john := j[day-1]
		
		a[day] = day - j[t_ann]
		j[day] = day - a[t_john]
		sum += a[day]
	}
	
	return sum
}

func sum_john(n int) int {
	if n == 0 {
		return 0
	}
	
	a := make([]int, n)
	j := make([]int, n)
	
	a[0] = 1
	j[0] = 0
	sum := 0
	
	for day := 1; day < n; day++ {
		t_ann := a[day-1]
		t_john := j[day-1]
		
		a[day] = day - j[t_ann]
		j[day] = day - a[t_john]
		sum += j[day]
	}
	
	return sum
}