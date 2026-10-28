package main

func chaserDistance(s int, t int) int {
    memo := make(map[[3]int]int)
    
    var solve func(timeLeft int, speed int, inRecovery bool) int
    solve = func(timeLeft int, speed int, inRecovery bool) int {
        if timeLeft == 0 || speed < 0 {
            return 0
        }
        
        recovery := 0
        if inRecovery {
            recovery = 1
        }
        
        key := [3]int{timeLeft, speed, recovery}
        if val, exists := memo[key]; exists {
            return val
        }
        
        var result int
        if inRecovery {
            result = speed + solve(timeLeft-1, speed, false)
        } else {
            runDistance := speed + solve(timeLeft-1, speed, false)
            sprintDistance := speed*2 + solve(timeLeft-1, speed-1, true)
            result = max(runDistance, sprintDistance)
        }
        
        memo[key] = result
        return result
    }
    
    return solve(t, s, false)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}