package kata

func AntBridge(ants string, terrain string) string {
	// find all gaps: each gap is a contiguous run of '.'
	type gap struct {
		size int
	}
	var gaps []gap
	i := 0
	for i < len(terrain) {
		if terrain[i] == '.' {
			j := i
			for j < len(terrain) && terrain[j] == '.' {
				j++
			}
			gaps = append(gaps, gap{size: j - i})
			i = j
		} else {
			i++
		}
	}

	// ants march left-to-right as a queue (index 0 is the front)
	queue := []byte(ants)

	for _, g := range gaps {
		bridgeSize := g.size

		// the first bridgeSize ants from the front form the bridge
		// they stop in order: first ant stops at edge, next climbs over, etc.
		// so bridge order from bottom to top (left to right in the gap) is
		// queue[0], queue[1], ..., queue[bridgeSize-1]
		bridge := make([]byte, bridgeSize)
		copy(bridge, queue[:bridgeSize])

		// remaining ants walk over the bridge
		remaining := queue[bridgeSize:]

		// after all remaining ants cross, the bridge dismantles:
		// the last ant placed on the bridge (topmost/rightmost) leaves first,
		// then next, etc. so they come off in reverse order of how they were placed
		// and join the back of the remaining ants queue

		// the ants that walked over are now on the other side in the same relative order
		// the bridge ants dismantle: last placed leaves first (rightmost first)
		// bridge[bridgeSize-1] leaves first, then bridge[bridgeSize-2], etc.
		// they rejoin at the back of the queue

		newQueue := make([]byte, 0, len(queue))
		newQueue = append(newQueue, remaining...)
		for j := bridgeSize - 1; j >= 0; j-- {
			newQueue = append(newQueue, bridge[j])
		}
		queue = newQueue
	}

	return string(queue)
}