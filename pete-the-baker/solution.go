package kata

func Cakes(recipe, available map[string]int) int {
	min := -1
	for ingredient, amount := range recipe {
		have, ok := available[ingredient]
		if !ok || amount <= 0 {
			return 0
		}
		n := have / amount
		if min == -1 || n < min {
			min = n
		}
	}
	if min == -1 {
		return 0
	}
	return min
}