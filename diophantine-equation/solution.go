package main

import (
	"sort"
)

func SolveEquation(n int) [][2]int {
	var solutions [][2]int
	
	for a := 1; a*a <= n; a++ {
		if n%a == 0 {
			b := n / a
			
			if a%4 == b%4 {
				x := (a + b) / 2
				y := (b - a) / 4
				solutions = append(solutions, [2]int{x, y})
			}
		}
	}
	
	sort.Slice(solutions, func(i, j int) bool {
		return solutions[i][0] > solutions[j][0]
	})
	
	return solutions
}