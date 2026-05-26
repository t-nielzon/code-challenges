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
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		code := line[len(line)-2:]
		full, ok := states[code]
		if !ok {
			continue
		}
		entry := strings.ReplaceAll(line, ",", "")
		entry = entry[:len(entry)-2] + full
		groups[full] = append(groups[full], entry)
	}

	names := make([]string, 0, len(groups))
	for k := range groups {
		names = append(names, k)
	}
	sort.Strings(names)

	var parts []string
	for i, st := range names {
		header := st
		if i > 0 {
			header = " " + st
		}
		parts = append(parts, header)
		sort.Strings(groups[st])
		for _, e := range groups[st] {
			parts = append(parts, "..... "+e)
		}
	}
	return strings.Join(parts, "\n")
}