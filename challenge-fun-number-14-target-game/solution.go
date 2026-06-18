package kata

// classic house-robber: at each target keep the best score achievable
// either by skipping it (prev) or shooting it plus the best from two back.
func TargetGame(vals []int) int {
	prev, prevPrev := 0, 0
	for _, v := range vals {
		take := prevPrev + v
		if take < prev {
			take = prev
		}
		prevPrev = prev
		prev = take
	}
	return prev
}