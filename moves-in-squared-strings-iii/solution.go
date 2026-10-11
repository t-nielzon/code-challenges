package main

import "strings"

func diag1Sym(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	
	result := make([]string, n)
	for col := 0; col < n; col++ {
		row := ""
		for i := 0; i < n; i++ {
			row += string(lines[i][col])
		}
		result[col] = row
	}
	return strings.Join(result, "\n")
}

func rot90Clock(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	
	result := make([]string, n)
	for row := 0; row < n; row++ {
		line := ""
		for col := n - 1; col >= 0; col-- {
			line += string(lines[col][row])
		}
		result[row] = line
	}
	return strings.Join(result, "\n")
}

func selfieAndDiag1(s string) string {
	lines := strings.Split(s, "\n")
	diagLines := strings.Split(diag1Sym(s), "\n")
	
	result := make([]string, len(lines))
	for i := 0; i < len(lines); i++ {
		result[i] = lines[i] + "|" + diagLines[i]
	}
	return strings.Join(result, "\n")
}

func oper(fct func(string) string, s string) string {
	return fct(s)
}