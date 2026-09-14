func chooseBestSum(t int, k int, ls []int) int {
	if len(ls) < k {
		return -1
	}

	maxSum := -1

	var combinations func(index int, count int, currentSum int)
	combinations = func(index int, count int, currentSum int) {
		if count == k {
			if currentSum <= t && currentSum > maxSum {
				maxSum = currentSum
			}
			return
		}

		if index >= len(ls) || count+len(ls)-index < k {
			return
		}

		// include current element
		combinations(index+1, count+1, currentSum+ls[index])

		// exclude current element
		combinations(index+1, count, currentSum)
	}

	combinations(0, 0, 0)

	return maxSum
}