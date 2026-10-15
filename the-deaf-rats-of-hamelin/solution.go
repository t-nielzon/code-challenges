package main

import "strings"

func CountDeafRats(s string) int {
	// Remove all spaces for uniform parsing
	s = strings.ReplaceAll(s, " ", "")
	
	// Find P position
	pIndex := strings.Index(s, "P")
	
	// Split by P
	left := s[:pIndex]
	right := s[pIndex+1:]
	
	deafCount := 0
	
	// Count ~O on the left (rats going right when they should go left)
	for i := 0; i < len(left); i += 2 {
		if i+1 < len(left) && left[i:i+2] == "~O" {
			deafCount++
		}
	}
	
	// Count O~ on the right (rats going left when they should go right)
	for i := 0; i < len(right); i += 2 {
		if i+1 < len(right) && right[i:i+2] == "O~" {
			deafCount++
		}
	}
	
	return deafCount
}