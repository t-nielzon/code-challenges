package main

import "math"

func cakes(recipe map[string]int, available map[string]int) int {
	minCakes := math.MaxInt

	for ingredient, required := range recipe {
		possibleCakes := available[ingredient] / required
		if possibleCakes < minCakes {
			minCakes = possibleCakes
		}
	}

	if minCakes == math.MaxInt {
		return 0
	}
	return minCakes
}