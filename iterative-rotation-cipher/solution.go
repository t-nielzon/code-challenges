package main

import (
	"fmt"
	"strings"
)

func Encode(n int, s string) string {
	result := s
	for i := 0; i < n; i++ {
		result = encodeIteration(n, result)
	}
	return fmt.Sprintf("%d %s", n, result)
}

func Decode(s string) string {
	parts := strings.SplitN(s, " ", 2)
	if len(parts) != 2 {
		return ""
	}
	
	var n int
	fmt.Sscanf(parts[0], "%d", &n)
	result := parts[1]
	
	for i := 0; i < n; i++ {
		result = decodeIteration(n, result)
	}
	return result
}

func encodeIteration(n int, s string) string {
	spacePositions := []int{}
	noSpaces := ""
	for i, ch := range s {
		if ch == ' ' {
			spacePositions = append(spacePositions, i)
		} else {
			noSpaces += string(ch)
		}
	}
	
	rotated := rotateRight(noSpaces, n)
	
	result := ""
	rotatedIdx := 0
	spaceIdx := 0
	for i := 0; i < len(s); i++ {
		if spaceIdx < len(spacePositions) && spacePositions[spaceIdx] == i {
			result += " "
			spaceIdx++
		} else {
			result += string(rotated[rotatedIdx])
			rotatedIdx++
		}
	}
	
	substrings := strings.Split(result, " ")
	for i := range substrings {
		substrings[i] = rotateRight(substrings[i], n)
	}
	
	return strings.Join(substrings, " ")
}

func decodeIteration(n int, s string) string {
	substrings := strings.Split(s, " ")
	for i := range substrings {
		substrings[i] = rotateLeft(substrings[i], n)
	}
	result := strings.Join(substrings, " ")
	
	spacePositions := []int{}
	noSpaces := ""
	for i, ch := range result {
		if ch == ' ' {
			spacePositions = append(spacePositions, i)
		} else {
			noSpaces += string(ch)
		}
	}
	
	rotated := rotateLeft(noSpaces, n)
	
	finalResult := ""
	rotatedIdx := 0
	spaceIdx := 0
	for i := 0; i < len(result); i++ {
		if spaceIdx < len(spacePositions) && spacePositions[spaceIdx] == i {
			finalResult += " "
			spaceIdx++
		} else {
			finalResult += string(rotated[rotatedIdx])
			rotatedIdx++
		}
	}
	
	return finalResult
}

func rotateRight(s string, n int) string {
	if len(s) == 0 {
		return s
	}
	n = n % len(s)
	return s[len(s)-n:] + s[:len(s)-n]
}

func rotateLeft(s string, n int) string {
	if len(s) == 0 {
		return s
	}
	n = n % len(s)
	return s[n:] + s[:n]
}