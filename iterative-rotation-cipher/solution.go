package main

import (
	"strconv"
	"strings"
)

func Encode(n int, text string) string {
	for i := 0; i < n; i++ {
		text = encodeOnce(n, text)
	}
	return strconv.Itoa(n) + " " + text
}

func encodeOnce(n int, text string) string {
	spacePositions := []int{}
	chars := []rune{}
	for i, ch := range text {
		if ch == ' ' {
			spacePositions = append(spacePositions, i)
		} else {
			chars = append(chars, ch)
		}
	}
	
	if len(chars) > 0 {
		n = n % len(chars)
		if n > 0 {
			rotated := make([]rune, len(chars))
			for i, ch := range chars {
				rotated[(i+n)%len(chars)] = ch
			}
			chars = rotated
		}
	}
	
	result := []rune{}
	charIdx := 0
	for i := 0; i < len(text); i++ {
		isSpace := false
		for _, pos := range spacePositions {
			if pos == i {
				isSpace = true
				break
			}
		}
		if isSpace {
			result = append(result, ' ')
		} else {
			result = append(result, chars[charIdx])
			charIdx++
		}
	}
	
	substrings := strings.Split(string(result), " ")
	for i := range substrings {
		substrings[i] = rotateRightStr(substrings[i], n)
	}
	result = []rune(strings.Join(substrings, " "))
	
	return string(result)
}

func rotateRightStr(s string, n int) string {
	if len(s) == 0 {
		return s
	}
	n = n % len(s)
	if n == 0 {
		return s
	}
	return s[len(s)-n:] + s[:len(s)-n]
}

func Decode(text string) string {
	parts := strings.SplitN(text, " ", 2)
	n, _ := strconv.Atoi(parts[0])
	encoded := parts[1]
	
	for i := 0; i < n; i++ {
		encoded = decodeOnce(n, encoded)
	}
	
	return encoded
}

func decodeOnce(n int, text string) string {
	substrings := strings.Split(text, " ")
	for i := range substrings {
		substrings[i] = rotateLeftStr(substrings[i], n)
	}
	text = strings.Join(substrings, " ")
	
	spacePositions := []int{}
	chars := []rune{}
	for i, ch := range text {
		if ch == ' ' {
			spacePositions = append(spacePositions, i)
		} else {
			chars = append(chars, ch)
		}
	}
	
	if len(chars) > 0 {
		n = n % len(chars)
		if n > 0 {
			rotated := make([]rune, len(chars))
			for i, ch := range chars {
				rotated[(i-n+len(chars))%len(chars)] = ch
			}
			chars = rotated
		}
	}
	
	result := []rune{}
	charIdx := 0
	for i := 0; i < len(text); i++ {
		isSpace := false
		for _, pos := range spacePositions {
			if pos == i {
				isSpace = true
				break
			}
		}
		if isSpace {
			result = append(result, ' ')
		} else {
			result = append(result, chars[charIdx])
			charIdx++
		}
	}
	
	return string(result)
}

func rotateLeftStr(s string, n int) string {
	if len(s) == 0 {
		return s
	}
	n = n % len(s)
	if n == 0 {
		return s
	}
	return s[n:] + s[:n]
}