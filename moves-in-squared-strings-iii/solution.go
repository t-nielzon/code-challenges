package main

import "strings"

func diag_1_sym(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	result := make([]string, n)
	
	for i := 0; i < n; i++ {
		row := ""
		for j := 0; j < n; j++ {
			row += string(lines[j][i])
		}
		result[i] = row
	}
	
	return strings.Join(result, "\n")
}

func rot_90_clock(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	result := make([]string, n)
	
	for i := 0; i < n; i++ {
		row := ""
		for j := 0; j < n; j++ {
			row += string(lines[n-1-j][i])
		}
		result[i] = row
	}
	
	return strings.Join(result, "\n")
}

func selfie_and_diag1(s string) string {
	diag := diag_1_sym(s)
	sLines := strings.Split(s, "\n")
	diagLines := strings.Split(diag, "\n")
	result := make([]string, len(sLines))
	
	for i := 0; i < len(sLines); i++ {
		result[i] = sLines[i] + "|" + diagLines[i]
	}
	
	return strings.Join(result, "\n")
}

func oper(fct func(string) string, s string) string {
	return fct(s)
}