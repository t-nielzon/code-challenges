package main

func chooseBestSum(t int, k int, ls []int) int {
	bestSum := -1

	var backtrack func(index int, remaining int, currentSum int)
	backtrack = func(index int, remaining int, currentSum int) {
		// prune if sum exceeds limit
		if currentSum > t {
			return
		}

		// found a valid combination of k elements
		if remaining == 0 {
			if currentSum > bestSum {
				bestSum = currentSum
			}
			return
		}

		// not enough elements left
		if index >= len(ls) {
			return
		}

		// include current element
		backtrack(index+1, remaining-1, currentSum+ls[index])

		// exclude current element
		backtrack(index+1, remaining, currentSum)
	}

	backtrack(0, k, 0)
	return bestSum
}