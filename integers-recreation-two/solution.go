package kata

import "sort"

func Prod2Sum(a, b, c, d int) [][2]int {
	// Using Brahmagupta–Fibonacci identity:
	// (a² + b²)(c² + d²) = (ac + bd)² + (ad - bc)² = (ac - bd)² + (ad + bc)²
	
	results := make(map[[2]int]bool)
	
	// First decomposition: (ac + bd, ad - bc)
	e1 := abs(a*c + b*d)
	f1 := abs(a*d - b*c)
	addPair(results, e1, f1)
	
	// Second decomposition: (ac - bd, ad + bc)
	e2 := abs(a*c - b*d)
	f2 := abs(a*d + b*c)
	addPair(results, e2, f2)
	
	// Convert map to slice
	var pairs [][2]int
	for pair := range results {
		pairs = append(pairs, pair)
	}
	
	// Sort by first element
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	
	return pairs
}

func addPair(results map[[2]int]bool, e, f int) {
	// Ensure lower element is first
	if e > f {
		e, f = f, e
	}
	results[[2]int{e, f}] = true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}