package main

import (
	"strconv"
	"strings"
)

func DrawRuler(t int, n int) string {
	lines := []string{}
	drawRulerHelper(&lines, t, n)
	return strings.Join(lines, "\n")
}

func drawRulerHelper(lines *[]string, t int, n int) {
	*lines = append(*lines, strings.Repeat("-", t)+" 0")
	
	for i := 1; i <= n; i++ {
		drawRulerSubdivisions(lines, t-1)
		*lines = append(*lines, strings.Repeat("-", t)+" "+strconv.Itoa(i))
	}
}

func drawRulerSubdivisions(lines *[]string, t int) {
	if t == 0 {
		return
	}
	
	drawRulerSubdivisions(lines, t-1)
	*lines = append(*lines, strings.Repeat("-", t))
	drawRulerSubdivisions(lines, t-1)
}