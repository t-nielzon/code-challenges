// solution.go

package main

func chasersSchedule(s, t int) int {
	memo := make(map[[3]int]int)

	var dp func(timeRemaining, speed int, inRecovery bool) int
	dp = func(timeRemaining, speed int, inRecovery bool) int {
		if timeRemaining == 0 {
			return 0
		}

		recovery := 0
		if inRecovery {
			recovery = 1
		}

		key := [3]int{timeRemaining, speed, recovery}
		if val, exists := memo[key]; exists {
			return val
		}

		var result int
		if inRecovery {
			// recovery is mandatory, no choice but to run at current speed
			result = speed + dp(timeRemaining-1, speed, false)
		} else {
			// choose sprint or run to maximize distance over remaining time
			runDist := speed + dp(timeRemaining-1, speed, false)
			sprintDist := speed*2 + dp(timeRemaining-1, speed-1, true)

			if sprintDist > runDist {
				result = sprintDist
			} else {
				result = runDist
			}
		}

		memo[key] = result
		return result
	}

	return dp(t, s, false)
}