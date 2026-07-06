package main

func QueueTime(customers []int, n int) int {
	if n == 0 || len(customers) == 0 {
		return 0
	}
	
	tills := make([]int, n)
	
	for _, time := range customers {
		minIdx := 0
		for i := 1; i < n; i++ {
			if tills[i] < tills[minIdx] {
				minIdx = i
			}
		}
		tills[minIdx] += time
	}
	
	maxTime := 0
	for _, t := range tills {
		if t > maxTime {
			maxTime = t
		}
	}
	
	return maxTime
}