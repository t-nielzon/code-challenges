package main

func cakes(recipe map[string]int, available map[string]int) int {
	var maxCakes int
	first := true
	
	for ingredient, needed := range recipe {
		have := available[ingredient]
		canMake := have / needed
		
		if first {
			maxCakes = canMake
			first = false
		} else if canMake < maxCakes {
			maxCakes = canMake
		}
	}
	
	return maxCakes
}