package kata

func QueueTime(customers []int, n int) int {
	if len(customers) == 0 {
		return 0
	}
	
	tills := make([]int, n)
	
	for _, customer := range customers {
		minIdx := 0
		for i := 1; i < n; i++ {
			if tills[i] < tills[minIdx] {
				minIdx = i
			}
		}
		tills[minIdx] += customer
	}
	
	maxTime := 0
	for _, till := range tills {
		if till > maxTime {
			maxTime = till
		}
	}
	
	return maxTime
}