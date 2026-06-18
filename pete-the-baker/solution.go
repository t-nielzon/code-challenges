package kata

func cakes(recipe, available map[string]int) int {
	min := -1
	for ingredient, needed := range recipe {
		have := available[ingredient]
		possible := have / needed
		if min == -1 || possible < min {
			min = possible
		}
	}
	if min == -1 {
		return 0
	}
	return min
}