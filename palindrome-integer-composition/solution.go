package main

import (
	"math"
	"strconv"
)

func values(n int) int {
	sumSet := make(map[int]bool)
	
	maxStart := int(math.Sqrt(float64(n)))
	for start := 1; start <= maxStart; start++ {
		sum := 0
		for current := start; ; current++ {
			sum += current * current
			if sum >= n {
				break
			}
			if current > start {
				sumSet[sum] = true
			}
		}
	}
	
	count := 0
	for num := range sumSet {
		if isPalindrome(num) {
			count++
		}
	}
	
	return count
}

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}