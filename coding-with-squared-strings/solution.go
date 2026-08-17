package main

import "strings"

func code(t string) string {
	if t == "" {
		return ""
	}
	
	l := len(t)
	n := 1
	for n*n < l {
		n++
	}
	
	padded := t
	for len(padded) < n*n {
		padded += string(rune(11))
	}
	
	square := make([][]rune, n)
	for i := 0; i < n; i++ {
		square[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			square[i][j] = rune(padded[i*n+j])
		}
	}
	
	rotated := make([][]rune, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]rune, n)
	}
	
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[j][n-1-i] = square[i][j]
		}
	}
	
	result := ""
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result += string(rotated[i][j])
		}
		if i < n-1 {
			result += "\n"
		}
	}
	
	return result
}

func decode(s string) string {
	if s == "" {
		return ""
	}
	
	lines := strings.Split(s, "\n")
	n := len(lines)
	
	square := make([][]rune, n)
	for i := 0; i < n; i++ {
		square[i] = []rune(lines[i])
	}
	
	rotated := make([][]rune, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]rune, n)
	}
	
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[n-1-j][i] = square[i][j]
		}
	}
	
	result := ""
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if rotated[i][j] != rune(11) {
				result += string(rotated[i][j])
			}
		}
	}
	
	return result
}