package main

import "strings"

func Diag2Sym(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	
	result := make([][]rune, n)
	for i := range result {
		result[i] = make([]rune, n)
	}
	
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[n-1-j][n-1-i] = rune(lines[i][j])
		}
	}
	
	return linesToString(result)
}

func Rot90Counter(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	
	result := make([][]rune, n)
	for i := range result {
		result[i] = make([]rune, n)
	}
	
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[n-1-j][i] = rune(lines[i][j])
		}
	}
	
	return linesToString(result)
}

func SelfieDiag2Counterclock(s string) string {
	diag2 := Diag2Sym(s)
	rot90 := Rot90Counter(s)
	
	origLines := strings.Split(s, "\n")
	diag2Lines := strings.Split(diag2, "\n")
	rot90Lines := strings.Split(rot90, "\n")
	
	var result []string
	for i := 0; i < len(origLines); i++ {
		result = append(result, origLines[i]+"|"+diag2Lines[i]+"|"+rot90Lines[i])
	}
	
	return strings.Join(result, "\n")
}

func linesToString(grid [][]rune) string {
	var lines []string
	for i := range grid {
		lines = append(lines, string(grid[i]))
	}
	return strings.Join(lines, "\n")
}

func Oper(fct func(string) string, s string) string {
	return fct(s)
}