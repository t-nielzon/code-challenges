package kata

import "strings"

type Operation func(string) string

func rot(s string) string {
	lines := strings.Split(s, "\n")
	var result []string
	
	for i := len(lines) - 1; i >= 0; i-- {
		reversed := reverseString(lines[i])
		result = append(result, reversed)
	}
	
	return strings.Join(result, "\n")
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func selfie_and_rot(s string) string {
	lines := strings.Split(s, "\n")
	if len(lines) == 0 || len(lines[0]) == 0 {
		return s
	}
	
	n := len(lines[0])
	dots := strings.Repeat(".", n)
	
	var result []string
	
	for _, line := range lines {
		result = append(result, line+dots)
	}
	
	rotated := rot(s)
	rotatedLines := strings.Split(rotated, "\n")
	for _, line := range rotatedLines {
		result = append(result, dots+line)
	}
	
	return strings.Join(result, "\n")
}

func oper(fct Operation, s string) string {
	return fct(s)
}