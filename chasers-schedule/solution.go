package main

func Distance(s, t int) int {
	memo := make(map[[2]int]int)
	
	var dp func(time, speed int) int
	dp = func(time, speed int) int {
		if time == 0 || speed <= 0 {
			return 0
		}
		
		key := [2]int{time, speed}
		if val, exists := memo[key]; exists {
			return val
		}
		
		// Option 1: run normally (don't sprint)
		result := speed + dp(time-1, speed)
		
		// Option 2: sprint
		if time == 1 {
			result = max(result, 2*speed)
		} else {
			// sprint at 2*speed, then forced recovery at (speed-1)
			sprintDist := 2*speed + (speed-1) + dp(time-2, speed-1)
			result = max(result, sprintDist)
		}
		
		memo[key] = result
		return result
	}
	
	return dp(t, s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}