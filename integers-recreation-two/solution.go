package kata

import "sort"

func Prod2Sum(a, b, c, d int) [][2]int {
	// brahmagupta–fibonacci identity: (a²+b²)(c²+d²) = (ac-bd)²+(ad+bc)² = (ac+bd)²+(ad-bc)²
	pairs := make(map[[2]int]bool)

	candidates := [][2]int{
		normalize(a*c+b*d, a*d-b*c),
		normalize(a*c-b*d, a*d+b*c),
	}

	for _, p := range candidates {
		pairs[p] = true
	}

	result := make([][2]int, 0, len(pairs))
	for p := range pairs {
		result = append(result, p)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

func normalize(x, y int) [2]int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if x > y {
		x, y = y, x
	}
	return [2]int{x, y}
}