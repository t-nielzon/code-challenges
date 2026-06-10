package kata

import "math"

func isSquare(x int) bool {
	r := int(math.Sqrt(float64(x)))
	return r*r == x || (r+1)*(r+1) == x
}

func SquareSumsRow(n int) []int {
	// adjacency: two numbers can be neighbors if their sum is a perfect square
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
		path = append(path, cur)
		used[cur] = true
		if len(path) == n {
			return true
		}
		for _, nxt := range adj[cur] {
			if !used[nxt] {
				if dfs(nxt) {
					return true
				}
			}
		}
		used[cur] = false
		path = path[:len(path)-1]
		return false
	}

	for start := 1; start <= n; start++ {
		path = path[:0]
		for i := range used {
			used[i] = false
		}
		if dfs(start) {
			result := make([]int, n)
			copy(result, path)
			return result
		}
	}
	return nil
}