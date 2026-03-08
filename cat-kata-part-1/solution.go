package kata

import "math"

func Peaceful(yard []string, minDistance int) bool {
	type pos struct{ r, c int }
	var cats []pos

	for r, row := range yard {
		for c, ch := range row {
			if ch == 'L' || ch == 'M' || ch == 'R' {
				cats = append(cats, pos{r, c})
			}
		}
	}

	if len(cats) <= 1 {
		return true
	}

	md := float64(minDistance)
	for i := 0; i < len(cats); i++ {
		for j := i + 1; j < len(cats); j++ {
			dr := float64(cats[i].r - cats[j].r)
			dc := float64(cats[i].c - cats[j].c)
			if math.Sqrt(dr*dr+dc*dc) < md {
				return false
			}
		}
	}
	return true
}