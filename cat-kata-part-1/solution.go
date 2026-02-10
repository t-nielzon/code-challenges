package kata

import "math"

func PeaceableKingdom(yard []string, minDistance int) bool {
	type point struct{ r, c int }
	var cats []point

	for r, row := range yard {
		for c, ch := range row {
			if ch == 'L' || ch == 'M' || ch == 'R' {
				cats = append(cats, point{r, c})
			}
		}
	}

	if len(cats) <= 1 {
		return true
	}

	for i := 0; i < len(cats); i++ {
		for j := i + 1; j < len(cats); j++ {
			dr := float64(cats[i].r - cats[j].r)
			dc := float64(cats[i].c - cats[j].c)
			dist := math.Sqrt(dr*dr + dc*dc)
			if dist < float64(minDistance) {
				return false
			}
		}
	}

	return true
}