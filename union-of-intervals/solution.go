package kata

func IntervalInsert(myl [][2]int, interval [2]int) [][2]int {
	result := [][2]int{}
	newInterval := interval
	i := 0

	// add all intervals that end before the new interval starts
	for i < len(myl) && myl[i][1] < newInterval[0] {
		result = append(result, myl[i])
		i++
	}

	// merge all overlapping intervals with the new interval
	for i < len(myl) && myl[i][0] <= newInterval[1] {
		if myl[i][0] < newInterval[0] {
			newInterval[0] = myl[i][0]
		}
		if myl[i][1] > newInterval[1] {
			newInterval[1] = myl[i][1]
		}
		i++
	}

	result = append(result, newInterval)

	// add remaining intervals
	for i < len(myl) {
		result = append(result, myl[i])
		i++
	}

	return result
}