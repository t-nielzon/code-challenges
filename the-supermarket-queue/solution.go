package main

func QueueTime(customers []int, n int) int {
	if len(customers) == 0 {
		return 0
	}
	
	// track the current time for each till
	tills := make([]int, n)
	
	// assign each customer to the till with minimum current time
	for _, time := range customers {
		// find the till with minimum time
		minIdx := 0
		for i := 1; i < n; i++ {
			if tills[i] < tills[minIdx] {
				minIdx = i
			}
		}
		// add customer to that till
		tills[minIdx] += time
	}
	
	// return the maximum time across all tills
	maxTime := 0
	for _, t := range tills {
		if t > maxTime {
			maxTime = t
		}
	}
	
	return maxTime
}