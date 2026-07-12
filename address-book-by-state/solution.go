package main

import (
	"fmt"
	"sort"
	"strings"
)

var stateMap = map[string]string{
	"AZ": "Arizona",
	"CA": "California",
	"ID": "Idaho",
	"IN": "Indiana",
	"MA": "Massachusetts",
	"OK": "Oklahoma",
	"PA": "Pennsylvania",
	"VA": "Virginia",
}

func AddressBookByState(str string) string {
	if str == "" {
		return ""
	}

	str = strings.ReplaceAll(str, "\r\n", "\n")
	lines := strings.Split(str, "\n")

	addresses := make(map[string][]map[string]string)

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
		address := strings.TrimSpace(parts[1])
		cityState := strings.TrimSpace(parts[2])

		cityStateParts := strings.Fields(cityState)
		if len(cityStateParts) < 2 {
			continue
		}

		stateCode := cityStateParts[len(cityStateParts)-1]
		city := strings.Join(cityStateParts[:len(cityStateParts)-1], " ")

		fullState, ok := stateMap[stateCode]
		if !ok {
			continue
		}

		person := map[string]string{
			"name":    name,
			"address": address,
			"city":    city,
			"state":   fullState,
		}

		addresses[fullState] = append(addresses[fullState], person)
	}

	var states []string
	for state := range addresses {
		states = append(states, state)
	}
	sort.Strings(states)

	var result []string
	for _, state := range states {
		result = append(result, state)

		people := addresses[state]
		sort.Slice(people, func(i, j int) bool {
			return people[i]["name"] < people[j]["name"]
		})

		for _, person := range people {
			line := fmt.Sprintf("..... %s %s %s %s",
				person["name"],
				person["address"],
				person["city"],
				person["state"],
			)
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}