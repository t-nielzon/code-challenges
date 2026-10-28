package kata

import (
	"sort"
	"strconv"
)

func MysteryRange(inputString string, size int) [2]int {
	var numbers []int
	
	var parse func(int, int) bool
	parse = func(index int, remaining int) bool {
		if remaining == 0 {
			return index == len(inputString)
		}
		
		if index >= len(inputString) {
			return false
		}
		
		for end := index + 1; end <= len(inputString); end++ {
			numStr := inputString[index:end]
			
			// skip numbers with leading zeros (except "0" itself)
			if len(numStr) > 1 && numStr[0] == '0' {
				continue
			}
			
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
			
			if parse(end, remaining-1) {
				return true
			}
			
			numbers = numbers[:len(numbers)-1]
		}
		
		return false
	}
	
	parse(0, size)
	
	sort.Ints(numbers)
	return [2]int{numbers[0], numbers[len(numbers)-1]}
}