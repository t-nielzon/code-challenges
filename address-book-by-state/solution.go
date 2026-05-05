package kata

import (
	"sort"
	"strings"
)

func AddressBook(s string) string {
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

	groups := map[string][]string{}
	for _, raw := range strings.Split(s, "\n") {
		line := strings.TrimRight(raw, "\r")
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Split(line, ",")
		last := strings.TrimSpace(parts[len(parts)-1])
		idx := strings.LastIndex(last, " ")
		if idx < 0 {
			continue
		}
		code := last[idx+1:]
		city := strings.TrimSpace(last[:idx])
		full, ok := states[code]
		if !ok {
			continue
		}
		var b strings.Builder
		for i := 0; i < len(parts)-1; i++ {
			b.WriteString(strings.TrimSpace(parts[i]))
			b.WriteString(" ")
		}
		b.WriteString(city)
		b.WriteString(" ")
		b.WriteString(full)
		groups[full] = append(groups[full], b.String())
	}

	stateNames := make([]string, 0, len(groups))
	for k := range groups {
		stateNames = append(stateNames, k)
	}
	sort.Strings(stateNames)

	var out []string
	for i, st := range stateNames {
		people := groups[st]
		sort.Strings(people)
		header := st
		if i > 0 {
			header = " " + st
		}
		block := header
		for _, p := range people {
			block += "\n..... " + p
		}
		out = append(out, block)
	}
	return strings.Join(out, "\n")
}