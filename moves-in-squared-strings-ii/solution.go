package main

import (
	"strings"
)

func Rot(s string) string {
	lines := strings.Split(s, "\n")
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
	for i := range lines {
		runes := []rune(lines[i])
		for j, k := 0, len(runes)-1; j < k; j, k = j+1, k-1 {
			runes[j], runes[k] = runes[k], runes[j]
		}
		lines[i] = string(runes)
	}
	return strings.Join(lines, "\n")
}

func SelfieAndRot(s string) string {
	lines := strings.Split(s, "\n")
	rotated := Rot(s)
	rotatedLines := strings.Split(rotated, "\n")
	
	dotCount := len(lines[0])
	dots := strings.Repeat(".", dotCount)
	
	result := []string{}
	for _, line := range lines {
		result = append(result, line+dots)
	}
	for _, line := range rotatedLines {
		result = append(result, dots+line)
	}
	return strings.Join(result, "\n")
}

func Oper(fct func(string) string, s string) string {
	return fct(s)
}