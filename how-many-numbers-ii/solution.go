package main

import (
    "math"
)

func MaxSumDig(nMax int, maxSum int) []int {
    var validNumbers []int
    var totalSum int64
    
    for num := 1000; num <= nMax; num++ {
        if isValid(num, maxSum) {
            validNumbers = append(validNumbers, num)
            totalSum += int64(num)
        }
    }
    
    count := len(validNumbers)
    if count == 0 {
        return []int{0, 0, 0}
    }
    
    mean := float64(totalSum) / float64(count)
    
    closest := validNumbers[0]
    minDist := math.Abs(float64(validNumbers[0]) - mean)
    
    for _, num := range validNumbers {
        dist := math.Abs(float64(num) - mean)
        
        if dist < minDist {
            minDist = dist
            closest = num
        } else if math.Abs(dist-minDist) < 1e-9 && num < closest {
            closest = num
        }
    }
    
    return []int{count, closest, int(totalSum)}
}

func isValid(num int, maxSum int) bool {
    digits := getDigits(num)
    
    for i := 0; i <= len(digits)-4; i++ {
        windowSum := 0
        for j := i; j < i+4; j++ {
            windowSum += digits[j]
        }
        if windowSum > maxSum {
            return false
        }
    }
    
    return true
}

func getDigits(num int) []int {
    var digits []int
    for num > 0 {
        digits = append([]int{num % 10}, digits...)
        num /= 10
    }
    return digits
}