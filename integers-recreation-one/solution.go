package kata

import "math"

func ListSquared(m, n uint64) [][2]uint64 {
	result := [][2]uint64{}
	for i := m; i <= n; i++ {
		var sum uint64 = 0
		sqrtI := uint64(math.Sqrt(float64(i)))
		for d := uint64(1); d <= sqrtI; d++ {
			if i%d == 0 {
				sum += d * d
				other := i / d
				if other != d {
					sum += other * other
				}
			}
		}
		s := uint64(math.Sqrt(float64(sum)))
		if s*s == sum {
			result = append(result, [2]uint64{i, sum})
		}
	}
	return result
}