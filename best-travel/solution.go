package main

func chooseBestSum(t int, k int, ls []int) int {
	result := -1
	
	var findCombinations func(index int, count int, sum int)
	findCombinations = func(index int, count int, sum int) {
		if count == k {
			if sum <= t && sum > result {
				result = sum
			}
			return
		}
		
		if index >= len(ls) || len(ls)-index < k-count {
			return
		}
		
		// include current distance
		findCombinations(index+1, count+1, sum+ls[index])
		
		// exclude current distance
		findCombinations(index+1, count, sum)
	}
	
	findCombinations(0, 0, 0)
	return result
}