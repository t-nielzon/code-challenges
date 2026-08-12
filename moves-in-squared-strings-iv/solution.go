package main

import "strings"

func diag_2_sym(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	
	result := make([][]byte, n)
	for i := 0; i < n; i++ {
		result[i] = make([]byte, n)
	}
	
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = lines[n-1-j][n-1-i]
		}
	}
	
	var res []string
	for i := 0; i < n; i++ {
		res = append(res, string(result[i]))
	}
	
	return strings.Join(res, "\n")
}

func rot_90_counter(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	
	result := make([][]byte, n)
	for i := 0; i < n; i++ {
		result[i] = make([]byte, n)
	}
	
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = lines[j][n-1-i]
		}
	}
	
	var res []string
	for i := 0; i < n; i++ {
		res = append(res, string(result[i]))
	}
	
	return strings.Join(res, "\n")
}

func selfie_diag2_counterclock(s string) string {
	lines := strings.Split(s, "\n")
	diag := diag_2_sym(s)
	rot := rot_90_counter(s)
	
	diagLines := strings.Split(diag, "\n")
	rotLines := strings.Split(rot, "\n")
	
	var result []string
	for i := 0; i < len(lines); i++ {
		combined := lines[i] + "|" + diagLines[i] + "|" + rotLines[i]
		result = append(result, combined)
	}
	
	return strings.Join(result, "\n")
}

func oper(fct func(string) string, s string) string {
	return fct(s)
}