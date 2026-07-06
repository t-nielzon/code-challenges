package kata

func QueueTime(customers []int, n int) int {
	if len(customers) == 0 {
		return 0
	}
	
	// track when each till will be free
	tills := make([]int, n)
	
	// assign each customer to the till that becomes free earliest
	for _, customer := range customers {
		// find the till with minimum free time
		minIdx := 0
		for i := 1; i < n; i++ {
			if tills[i] < tills[minIdx] {
				minIdx = i
			}
		}
		
		// add customer to that till
		tills[minIdx] += customer
	}
	
	// return the maximum time (when the last till finishes)
	maxTime := tills[0]
	for i := 1; i < n; i++ {
		if tills[i] > maxTime {
			maxTime = tills[i]
		}
	}
	
	return maxTime
}