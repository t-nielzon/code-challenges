package kata

import (
	"sort"
	"strings"
)

func TravelInfos(r string) string {
	states := map[string]string{
		"AZ": "Arizona",
		"CA": "California",
		"ID": "Idaho",
		"IN": "Indiana",
		"MA": "Massachusetts",
		"OK": "Oklahoma",
		"PA": "Pennsylvania",
		"VA": "Virginia",
	}

	grouped := map[string][]string{}
	for _, line := range strings.Split(r, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ", ")
		if len(parts) < 3 {
			continue
		}
		name := parts[0]
		street := parts[1]
		cityState := parts[2]
		idx := strings.LastIndex(cityState, " ")
		if idx < 0 {
			continue
		}
		city := cityState[:idx]
		code := cityState[idx+1:]
		fullState, ok := states[code]
		if !ok {
			continue
		}
		entry := "..... " + name + " " + street + " " + city + " " + fullState
		grouped[fullState] = append(grouped[fullState], entry)
	}

	stateNames := make([]string, 0, len(grouped))
	for s := range grouped {
		stateNames = append(stateNames, s)
	}
	sort.Strings(stateNames)

	blocks := make([]string, 0, len(stateNames))
	for _, s := range stateNames {
		entries := grouped[s]
		sort.Strings(entries)
		blocks = append(blocks, s+"\n"+strings.Join(entries, "\n"))
	}
	return strings.Join(blocks, "\n ")
}