package solution

import (
	"strings"
)

func Closest(strng string) [][]interface{} {
	if strng == "" {
		return [][]interface{}{}
	}
	
	parts := strings.Fields(strng)
	if len(parts) == 0 {
		return [][]interface{}{}
	}
	
	type NumInfo struct {
		weight int
		index  int
		num    string
	}
	
	nums := make([]NumInfo, len(parts))
	for i, part := range parts {
		weight := 0
		for _, ch := range part {
			weight += int(ch - '0')
		}
		nums[i] = NumInfo{weight: weight, index: i, num: part}
	}
	
	var bestPair [2]NumInfo
	bestDiff := 1<<31 - 1
	bestMinWeight := 1<<31 - 1
	bestMinIndex := 1<<31 - 1
	
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			diff := abs(nums[i].weight - nums[j].weight)
			minWeight := min(nums[i].weight, nums[j].weight)
			minIndex := min(nums[i].index, nums[j].index)
			
			if diff < bestDiff ||
				(diff == bestDiff && minWeight < bestMinWeight) ||
				(diff == bestDiff && minWeight == bestMinWeight && minIndex < bestMinIndex) {
				bestDiff = diff
				bestMinWeight = minWeight
				bestMinIndex = minIndex
				bestPair = [2]NumInfo{nums[i], nums[j]}
			}
		}
	}
	
	first := bestPair[0]
	second := bestPair[1]
	if first.weight > second.weight || (first.weight == second.weight && first.index > second.index) {
		first, second = second, first
	}
	
	return [][]interface{}{
		{first.weight, first.index, first.num},
		{second.weight, second.index, second.num},
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}