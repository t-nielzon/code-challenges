package kata

import "math"

func SquareSumsRow(n int) []int {
	// build adjacency list: for each number 1..n, find valid neighbors
	adj := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			s := i + j
			sqrt := int(math.Round(math.Sqrt(float64(s))))
			if sqrt*sqrt == s {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}

	path := make([]int, 0, n)
	used := make([]bool, n+1)

	var backtrack func() bool
	backtrack = func() bool {
		if len(path) == n {
			return true
		}
		var candidates []int
		if len(path) == 0 {
			// try all starting numbers
			candidates = make([]int, n)
			for i := range candidates {
				candidates[i] = i + 1
			}
		} else {
			candidates = adj[path[len(path)-1]]
		}
		for _, next := range candidates {
			if !used[next] {
				used[next] = true
				path = append(path, next)
				if backtrack() {
					return true
				}
				path = path[:len(path)-1]
				used[next] = false
			}
		}
		return false
	}

	if backtrack() {
		result := make([]int, n)
		copy(result, path)
		return result
	}
	return nil
}