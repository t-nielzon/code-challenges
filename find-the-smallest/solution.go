package main

import (
	"strconv"
)

func smallest(n int) [3]int {
	s := strconv.Itoa(n)
	digits := []rune(s)
	minNum := n
	minI, minJ := 0, 0
	
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits); j++ {
			temp := make([]rune, len(digits))
			copy(temp, digits)
			
			digit := temp[i]
			temp = append(temp[:i], temp[i+1:]...)
			
			temp = append(temp[:j], append([]rune{digit}, temp[j:]...)...)
			
			numStr := string(temp)
			num, _ := strconv.Atoi(numStr)
			
			if num < minNum || (num == minNum && i < minI) || (num == minNum && i == minI && j < minJ) {
				minNum = num
				minI = i
				minJ = j
			}
		}
	}
	
	return [3]int{minNum, minI, minJ}
}