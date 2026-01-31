package kata

func Cakes(recipe, available map[string]int) int {
	if len(recipe) == 0 {
		return 0
	}

	minCakes := -1

	for ingredient, needed := range recipe {
		if needed == 0 {
			continue
		}

		have, exists := available[ingredient]
		if !exists {
			return 0
		}

		possible := have / needed

		if minCakes == -1 || possible < minCakes {
			minCakes = possible
		}
	}

	if minCakes == -1 {
		return 0
	}

	return minCakes
}