package main

func TwoSum(numbers []int, target int) [2]int {
	seen := make(map[int]int)
	
	for i, num := range numbers {
		complement := target - num
		if j, exists := seen[complement]; exists {
			return [2]int{j, i}
		}
		seen[num] = i
	}
	
	return [2]int{}
}