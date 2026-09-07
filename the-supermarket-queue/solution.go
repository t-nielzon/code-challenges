package main

func QueueTime(customers []int, n int) int {
	if n == 0 || len(customers) == 0 {
		return 0
	}

	// tills tracks when each till will be free
	tills := make([]int, n)

	for _, customer := range customers {
		// find the till that will be free earliest
		minTill := 0
		for i := 1; i < n; i++ {
			if tills[i] < tills[minTill] {
				minTill = i
			}
		}

		// assign customer to this till
		tills[minTill] += customer
	}

	// return the maximum time (when the last till finishes)
	maxTime := 0
	for _, till := range tills {
		if till > maxTime {
			maxTime = till
		}
	}

	return maxTime
}