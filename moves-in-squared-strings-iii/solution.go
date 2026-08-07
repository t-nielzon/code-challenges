package kata

import (
	"strings"
)

func diag_1_sym(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	
	var result []string
	for i := 0; i < n; i++ {
		var row string
		for j := 0; j < n; j++ {
			row += string(lines[j][i])
		}
		result = append(result, row)
	}
	
	return strings.Join(result, "\n")
}

func rot_90_clock(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	
	var result []string
	for i := 0; i < n; i++ {
		var row string
		for j := n - 1; j >= 0; j-- {
			row += string(lines[j][i])
		}
		result = append(result, row)
	}
	
	return strings.Join(result, "\n")
}

func selfie_and_diag1(s string) string {
	diag := diag_1_sym(s)
	lines := strings.Split(s, "\n")
	diagLines := strings.Split(diag, "\n")
	
	var result []string
	for i := 0; i < len(lines); i++ {
		result = append(result, lines[i]+"|"+diagLines[i])
	}
	
	return strings.Join(result, "\n")
}

func oper(fct func(string) string, s string) string {
	return fct(s)
}