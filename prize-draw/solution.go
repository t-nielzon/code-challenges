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

	var participants []participant

	for i, name := range names {
		som := len(name)
		for _, char := range name {
			if char >= 'A' && char <= 'Z' {
				som += int(char - 'A' + 1)
			} else if char >= 'a' && char <= 'z' {
				som += int(char - 'a' + 1)
			}
		}
		winningNumber := som * we[i]
		participants = append(participants, participant{name, winningNumber})
	}

	sort.Slice(participants, func(i, j int) bool {
		if participants[i].winningNumber != participants[j].winningNumber {
			return participants[i].winningNumber > participants[j].winningNumber
		}
		return participants[i].name < participants[j].name
	})

	return participants[n-1].name
}