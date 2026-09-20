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

	type Participant struct {
		name          string
		winningNumber int
	}

	participants := make([]Participant, len(names))

	for i, name := range names {
		som := len(name)
		for _, ch := range name {
			if ch >= 'A' && ch <= 'Z' {
				som += int(ch - 'A' + 1)
			} else if ch >= 'a' && ch <= 'z' {
				som += int(ch - 'a' + 1)
			}
		}

		winningNumber := som * we[i]

		participants[i] = Participant{
			name:          name,
			winningNumber: winningNumber,
		}
	}

	sort.Slice(participants, func(i, j int) bool {
		if participants[i].winningNumber != participants[j].winningNumber {
			return participants[i].winningNumber > participants[j].winningNumber
		}
		return participants[i].name < participants[j].name
	})

	return participants[n-1].name
}