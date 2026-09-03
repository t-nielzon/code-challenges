package main

func cakes(recipe map[string]int, available map[string]int) int {
	minCakes := -1
	
	for ingredient, needed := range recipe {
		have := available[ingredient]
		canMake := have / needed
		
		if minCakes == -1 || canMake < minCakes {
			minCakes = canMake
		}
	}
	
	if minCakes == -1 {
		return 0
	}
	
	return minCakes
}