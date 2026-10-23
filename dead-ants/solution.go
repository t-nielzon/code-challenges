package main

import "strings"

func DeadAnts(s string) int {
	// Remove all complete "ant" substrings (these ants survived)
	for strings.Contains(s, "ant") {
		s = strings.Replace(s, "ant", "", 1)
	}

	// Count remaining body parts from dead ants
	a := strings.Count(s, "a")
	n := strings.Count(s, "n")
	t := strings.Count(s, "t")

	// Each dead ant needs at least one of each character (a, n, t)
	// The minimum count determines how many complete ants we can form
	min := a
	if n < min {
		min = n
	}
	if t < min {
		min = t
	}
	return min
}