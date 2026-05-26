package kata

import "math"

func isSquare(x int) bool {
	s := int(math.Sqrt(float64(x)))
	return s*s == x || (s+1)*(s+1) == x
}

func SquareSums(n int) []int {
	used := make([]bool, n+1)
	path := make([]int, 0, n)

	var backtrack func() bool
	backtrack = func() bool {
		if len(path) == n {
			return true
		}
		for i := 1; i <= n; i++ {
			if used[i] {
				continue
			}
			if len(path) > 0 && !isSquare(path[len(path)-1]+i) {
				continue
			}
			used[i] = true
			path = append(path, i)
			if backtrack() {
				return true
			}
			path = path[:len(path)-1]
			used[i] = false
		}
		return false
	}

	if backtrack() {
		result := make([]int, len(path))
		copy(result, path)
		return result
	}
	return nil
}