package main

func chaserScore(s, t int) int {
	memo := make(map[[3]int]int)
	
	var dp func(timeLeft, speed int, inRecovery bool) int
	dp = func(timeLeft, speed int, inRecovery bool) int {
		if timeLeft == 0 || speed <= 0 {
			return 0
		}
		
		key := [3]int{timeLeft, speed, 0}
		if inRecovery {
			key[2] = 1
		}
		
		if cached, exists := memo[key]; exists {
			return cached
		}
		
		var result int
		if inRecovery {
			// recovery phase: must run normally
			result = speed + dp(timeLeft-1, speed, false)
		} else {
			// normal phase: choose between running normally or sprinting
			runNormally := speed + dp(timeLeft-1, speed, false)
			sprint := speed*2 + dp(timeLeft-1, speed-1, true)
			
			result = runNormally
			if sprint > runNormally {
				result = sprint
			}
		}
		
		memo[key] = result
		return result
	}
	
	return dp(t, s, false)
}