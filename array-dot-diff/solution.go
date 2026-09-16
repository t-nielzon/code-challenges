package main

func ArrayDiff(a, b []int) []int {
	bMap := make(map[int]bool)
	for _, v := range b {
		bMap[v] = true
	}
	
	var result []int
	for _, v := range a {
		if !bMap[v] {
			result = append(result, v)
		}
	}
	
	return result
}