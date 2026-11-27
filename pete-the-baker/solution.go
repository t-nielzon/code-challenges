package kata

func Cakes(recipe map[string]int, available map[string]int) int {
	minCakes := -1

	for ingredient, needed := range recipe {
		cakesFromThis := available[ingredient] / needed
		if minCakes == -1 || cakesFromThis < minCakes {
			minCakes = cakesFromThis
		}
	}

	if minCakes < 0 {
		return 0
	}
	return minCakes
}