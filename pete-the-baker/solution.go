package kata

func Cakes(recipe, available map[string]int) int {
	min := -1
	for ingredient, needed := range recipe {
		have := available[ingredient]
		count := have / needed
		if min == -1 || count < min {
			min = count
		}
	}
	if min == -1 {
		return 0
	}
	return min
}