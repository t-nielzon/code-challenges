package main

import (
	"sort"
	"strings"
)

func PrizeDraw(st string, we []int, n int) string {
	if st == "" {
		return "No participants"
	}
	
	names := strings.Split(st, ",")
	
	if n > len(names) {
		return "Not enough participants"
	}
	
	type participant struct {
		name          string
		winningNumber int
	}
	
	participants := make([]participant, len(names))
	
	for i, name := range names {
		som := len(name)
		for _, ch := range name {
			var rank int
			if ch >= 'a' && ch <= 'z' {
				rank = int(ch - 'a' + 1)
			} else if ch >= 'A' && ch <= 'Z' {
				rank = int(ch - 'A' + 1)
			}
			som += rank
		}
		
		winningNumber := som * we[i]
		participants[i] = participant{name, winningNumber}
	}
	
	sort.Slice(participants, func(i, j int) bool {
		if participants[i].winningNumber != participants[j].winningNumber {
			return participants[i].winningNumber > participants[j].winningNumber
		}
		return participants[i].name < participants[j].name
	})
	
	return participants[n-1].name
}