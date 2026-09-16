package main

func cakes(recipe, available map[string]int) int {
	minCakes := -1

	for ingredient, required := range recipe {
		possibleCakes := available[ingredient] / required
		if minCakes == -1 || possibleCakes < minCakes {
			minCakes = possibleCakes
		}
	}

	if minCakes == -1 {
		return 0
	}

	return minCakes
}