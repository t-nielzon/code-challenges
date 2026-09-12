package main

import (
	"sort"
	"strings"
)

func AddressBookByState(address string) string {
	stateMap := map[string]string{
		"AZ": "Arizona",
		"CA": "California",
		"ID": "Idaho",
		"IN": "Indiana",
		"MA": "Massachusetts",
		"OK": "Oklahoma",
		"PA": "Pennsylvania",
		"VA": "Virginia",
	}

	type Person struct {
		name     string
		fullLine string
	}

	stateGroups := make(map[string][]Person)

	lines := strings.Split(address, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) < 3 {
			continue
		}

		name := strings.TrimSpace(parts[0])
		streetAddr := strings.TrimSpace(parts[1])
		cityStateStr := strings.TrimSpace(parts[2])

		words := strings.Fields(cityStateStr)
		if len(words) < 2 {
			continue
		}

		stateCode := words[len(words)-1]
		city := strings.Join(words[:len(words)-1], " ")

		fullStateName := stateMap[stateCode]
		if fullStateName == "" {
			continue
		}

		fullLine := name + " " + streetAddr + " " + city + " " + fullStateName

		person := Person{
			name:     name,
			fullLine: fullLine,
		}

		stateGroups[fullStateName] = append(stateGroups[fullStateName], person)
	}

	for state := range stateGroups {
		sort.Slice(stateGroups[state], func(i, j int) bool {
			return stateGroups[state][i].name < stateGroups[state][j].name
		})
	}

	var states []string
	for state := range stateGroups {
		states = append(states, state)
	}
	sort.Strings(states)

	var result []string
	for _, state := range states {
		result = append(result, state)
		for _, person := range stateGroups[state] {
			result = append(result, "..... "+person.fullLine)
		}
	}

	return strings.Join(result, "\n")
}