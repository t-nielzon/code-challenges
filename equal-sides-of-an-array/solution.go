// solution.go
package main

func FindEvenIndex(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	
	leftSum := 0
	for i, v := range arr {
		rightSum := total - leftSum - v
		if leftSum == rightSum {
			return i
		}
		leftSum += v
	}
	
	return -1
}