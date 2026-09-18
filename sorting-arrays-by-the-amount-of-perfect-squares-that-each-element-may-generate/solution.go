package main

import (
	"math"
	"sort"
	"strconv"
)

func SortByPerfsq(arr []int) []int {
	cache := make(map[int]int)
	
	for _, num := range arr {
		cache[num] = countPerfectSquares(num)
	}
	
	sort.Slice(arr, func(i, j int) bool {
		countI := cache[arr[i]]
		countJ := cache[arr[j]]
		
		if countI != countJ {
			return countI > countJ
		}
		return arr[i] < arr[j]
	})
	
	return arr
}

func countPerfectSquares(num int) int {
	digits := strconv.Itoa(num)
	perms := getPermutations(digits)
	
	perfSqs := make(map[int]bool)
	for _, perm := range perms {
		val, _ := strconv.Atoi(perm)
		if isPerfectSquare(val) {
			perfSqs[val] = true
		}
	}
	
	return len(perfSqs)
}

func getPermutations(s string) []string {
	var result []string
	permute([]rune(s), 0, len(s)-1, &result)
	
	seen := make(map[string]bool)
	var unique []string
	for _, perm := range result {
		if !seen[perm] {
			seen[perm] = true
			unique = append(unique, perm)
		}
	}
	
	return unique
}

func permute(runes []rune, l, r int, result *[]string) {
	if l == r {
		*result = append(*result, string(runes))
	} else {
		for i := l; i <= r; i++ {
			runes[l], runes[i] = runes[i], runes[l]
			permute(runes, l+1, r, result)
			runes[l], runes[i] = runes[i], runes[l]
		}
	}
}

func isPerfectSquare(num int) bool {
	sqrt := int(math.Sqrt(float64(num)))
	return sqrt*sqrt == num
}