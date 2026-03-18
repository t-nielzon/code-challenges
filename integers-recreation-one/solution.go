package kata

import "math"

func ListSquared(m, n int) [][]int {
	var result [][]int
	for i := m; i <= n; i++ {
		sum := 0
		for d := 1; d*d <= i; d++ {
			if i%d == 0 {
				sum += d * d
				if d != i/d {
					sum += (i / d) * (i / d)
				}
			}
		}
		sqrt := int(math.Sqrt(float64(sum)))
		if sqrt*sqrt == sum {
			result = append(result, []int{i, sum})
		}
	}
	return result
}