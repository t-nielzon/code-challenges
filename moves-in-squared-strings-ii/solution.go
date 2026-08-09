package main

import "strings"

func rot(s string) string {
	lines := strings.Split(s, "\n")
	
	// reverse order of lines and reverse each line for 180 degree rotation
	for i := 0; i < len(lines)/2; i++ {
		j := len(lines) - 1 - i
		lines[i], lines[j] = lines[j], lines[i]
	}
	
	for i := range lines {
		runes := []rune(lines[i])
		for j := 0; j < len(runes)/2; j++ {
			k := len(runes) - 1 - j
			runes[j], runes[k] = runes[k], runes[j]
		}
		lines[i] = string(runes)
	}
	
	return strings.Join(lines, "\n")
}

func selfie_and_rot(s string) string {
	if s == "" {
		return ""
	}
	
	lines := strings.Split(s, "\n")
	n := len(lines[0])
	
	rotated := rot(s)
	rotLines := strings.Split(rotated, "\n")
	
	var result []string
	
	for _, line := range lines {
		result = append(result, line+strings.Repeat(".", n))
	}
	
	for _, line := range rotLines {
		result = append(result, strings.Repeat(".", n)+line)
	}
	
	return strings.Join(result, "\n")
}

func oper(fct func(string) string, s string) string {
	return fct(s)
}