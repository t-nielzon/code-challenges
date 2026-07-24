package main

import (
	"strconv"
)

func smallest(n int) [3]int {
	s := strconv.Itoa(n)
	digits := []rune(s)
	
	var result [3]int
	result[0] = n
	result[1] = 0
	result[2] = 0
	
	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		
		temp := append([]rune{}, digits[:i]...)
		temp = append(temp, digits[i+1:]...)
		
		for j := 0; j <= len(temp); j++ {
			newDigits := append([]rune{}, temp[:j]...)
			newDigits = append(newDigits, digit)
			newDigits = append(newDigits, temp[j:]...)
			
			num, _ := strconv.Atoi(string(newDigits))
			
			if num < result[0] || (num == result[0] && i < result[1]) || (num == result[0] && i == result[1] && j < result[2]) {
				result[0] = num
				result[1] = i
				result[2] = j
			}
		}
	}
	
	return result
}