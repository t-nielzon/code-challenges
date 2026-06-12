package kata

import "math"

func isSquare(x int) bool {
	r := int(math.Sqrt(float64(x)))
	return r*r == x
}

func SquareSums(n int) []int {
	// adjacency: numbers whose pairwise sum is a perfect square
	adj := make([][]int, n+1)
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			if a != b && isSquare(a+b) {
				adj[a] = append(adj[a], b)
			}
		}
	}

	used := make([]bool, n+1)
	path := make([]int, 0, n)

	var dfs func(cur int) bool
	dfs = func(cur int) bool {
		used[cur] = true
		path = append(path, cur)
		if len(path) == n {
			return true
		}
		// prefer neighbors with fewest remaining options to prune the
		// search early (warnsdorff-style heuristic keeps n<=43 fast)
		candidates := []int{}
		for _, nxt := range adj[cur] {
			if !used[nxt] {
				candidates = append(candidates, nxt)
			}
		}
		degree := func(v int) int {
			d := 0
			for _, w := range adj[v] {
				if !used[w] {
					d++
				}
			}
			return d
		}
		for i := 1; i < len(candidates); i++ {
			for j := i; j > 0 && degree(candidates[j]) < degree(candidates[j-1]); j-- {
				candidates[j], candidates[j-1] = candidates[j-1], candidates[j]
			}
		}
		for _, nxt := range candidates {
			if dfs(nxt) {
				return true
			}
		}
		used[cur] = false
		path = path[:len(path)-1]
		return false
	}

	for start := 1; start <= n; start++ {
		if dfs(start) {
			return path
		}
	}
	return nil
}