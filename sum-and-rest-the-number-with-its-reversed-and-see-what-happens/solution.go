package main

var memo = make(map[int]int)
var currentIndex = 0
var currentNum = 0

func reverseNumber(num int) int {
	reversed := 0
	for num > 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	return reversed
}

func hasSpecialProperty(num int) bool {
	if num%10 == 0 {
		return false
	}
	
	reversed := reverseNumber(num)
	sum := num + reversed
	diff := num - reversed
	if diff < 0 {
		diff = -diff
	}
	
	if diff == 0 {
		return false
	}
	
	return sum%diff == 0
}

func Nthterm(n int) int {
	if val, ok := memo[n]; ok {
		return val
	}
	
	for currentIndex < n {
		currentNum++
		if hasSpecialProperty(currentNum) {
			currentIndex++
			memo[currentIndex] = currentNum
		}
	}
	
	return memo[n]
}