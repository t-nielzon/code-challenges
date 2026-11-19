package main

import "sort"

func Prod2sum(a, b, c, d int) [][]int {
	// using Brahmagupta–Fibonacci identity:
	// (a² + b²)(c² + d²) = (ac - bd)² + (ad + bc)²
	//                    = (ac + bd)² + (ad - bc)²
	
	ac := a * c
	ad := a * d
	bc := b * c
	bd := b * d
	
	// first decomposition
	e1 := abs(ac - bd)
	f1 := abs(ad + bc)
	
	// second decomposition
	e2 := abs(ac + bd)
	f2 := abs(ad - bc)
	
	// create pairs with smaller element first
	var pairs [][]int
	
	pair1 := []int{minInt(e1, f1), maxInt(e1, f1)}
	pair2 := []int{minInt(e2, f2), maxInt(e2, f2)}
	
	// add pairs if they're different (remove duplicates)
	pairs = append(pairs, pair1)
	if !arraysEqual(pair1, pair2) {
		pairs = append(pairs, pair2)
	}
	
	// sort pairs by first element
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	
	return pairs
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func arraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}