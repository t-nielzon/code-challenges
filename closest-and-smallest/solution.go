package main

import (
	"strconv"
	"strings"
)

func Closest(strng string) [][]interface{} {
	if strng == "" {
		return [][]interface{}{}
	}

	parts := strings.Fields(strng)
	if len(parts) < 2 {
		return [][]interface{}{}
	}

	type NumInfo struct {
		weight   int
		index    int
		original int
	}

	nums := make([]NumInfo, len(parts))
	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		nums[i] = NumInfo{
			weight:   calculateWeight(num),
			index:    i,
			original: num,
		}
	}

	var bestI, bestJ int = -1, -1
	minDiff := -1
	var bestWeights [2]int
	bestMaxIndex := -1
	bestMinIndex := -1

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			diff := abs(nums[i].weight - nums[j].weight)
			w1, w2 := nums[i].weight, nums[j].weight
			if w1 > w2 {
				w1, w2 = w2, w1
			}
			maxIdx := max(nums[i].index, nums[j].index)
			minIdx := min(nums[i].index, nums[j].index)

			isBetter := false
			if bestI == -1 {
				isBetter = true
			} else if diff < minDiff {
				isBetter = true
			} else if diff == minDiff {
				if w1 < bestWeights[0] || (w1 == bestWeights[0] && w2 < bestWeights[1]) {
					isBetter = true
				} else if w1 == bestWeights[0] && w2 == bestWeights[1] {
					if maxIdx < bestMaxIndex {
						isBetter = true
					} else if maxIdx == bestMaxIndex && minIdx < bestMinIndex {
						isBetter = true
					}
				}
			}

			if isBetter {
				minDiff = diff
				bestWeights = [2]int{w1, w2}
				bestMaxIndex = maxIdx
				bestMinIndex = minIdx
				bestI = i
				bestJ = j
			}
		}
	}

	first := nums[bestI]
	second := nums[bestJ]

	if first.weight > second.weight || 
		(first.weight == second.weight && first.index > second.index) {
		first, second = second, first
	}

	return [][]interface{}{
		{first.weight, first.index, first.original},
		{second.weight, second.index, second.original},
	}
}

func calculateWeight(num int) int {
	weight := 0
	for num > 0 {
		weight += num % 10
		num /= 10
	}
	return weight
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}