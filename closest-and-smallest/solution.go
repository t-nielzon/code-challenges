package main

import (
	"strconv"
	"strings"
)

func closest(strng string) [][]int {
	if strng == "" {
		return [][]int{}
	}

	parts := strings.Fields(strng)
	if len(parts) < 2 {
		return [][]int{}
	}

	weights := make([]int, len(parts))
	numbers := make([]int, len(parts))
	for i, part := range parts {
		weight := 0
		for _, ch := range part {
			weight += int(ch - '0')
		}
		weights[i] = weight
		numbers[i], _ = strconv.Atoi(part)
	}

	var bestI, bestJ int
	var bestDiff int = -1
	var bestSum int

	for i := 0; i < len(weights); i++ {
		for j := i + 1; j < len(weights); j++ {
			diff := abs(weights[i] - weights[j])
			sum := weights[i] + weights[j]

			if bestDiff == -1 ||
				diff < bestDiff ||
				(diff == bestDiff && sum < bestSum) ||
				(diff == bestDiff && sum == bestSum && i < bestI) ||
				(diff == bestDiff && sum == bestSum && i == bestI && j < bestJ) {
				bestI, bestJ = i, j
				bestDiff = diff
				bestSum = sum
			}
		}
	}

	arr1 := []int{weights[bestI], bestI, numbers[bestI]}
	arr2 := []int{weights[bestJ], bestJ, numbers[bestJ]}

	if arr1[0] > arr2[0] || (arr1[0] == arr2[0] && arr1[1] > arr2[1]) {
		arr1, arr2 = arr2, arr1
	}

	return [][]int{arr1, arr2}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}