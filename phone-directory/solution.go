package main

import (
	"fmt"
	"regexp"
	"strings"
)

func phone(s string, num string) string {
	lines := strings.Split(s, "\n")

	var matches []map[string]string

	phoneRegex := regexp.MustCompile(`\+(\d{1,2}-[\d\-]+)`)
	nameRegex := regexp.MustCompile(`<([^>]+)>`)

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Extract phone number
		phoneMatches := phoneRegex.FindStringSubmatch(line)
		if phoneMatches == nil {
			continue
		}

		phoneNum := phoneMatches[1]
		if phoneNum != num {
			continue
		}

		// Extract name
		nameMatches := nameRegex.FindStringSubmatch(line)
		if nameMatches == nil {
			continue
		}
		name := strings.TrimSpace(nameMatches[1])

		// Extract and clean address
		address := line
		address = phoneRegex.ReplaceAllString(address, " ")
		address = nameRegex.ReplaceAllString(address, " ")
		address = regexp.MustCompile(`[^a-zA-Z0-9\s.]`).ReplaceAllString(address, " ")
		address = regexp.MustCompile(`\s+`).ReplaceAllString(address, " ")
		address = strings.TrimSpace(address)

		matches = append(matches, map[string]string{
			"phone":   num,
			"name":    name,
			"address": address,
		})
	}

	if len(matches) == 0 {
		return fmt.Sprintf("Error => Not found: %s", num)
	}
	if len(matches) > 1 {
		return fmt.Sprintf("Error => Too many people: %s", num)
	}

	m := matches[0]
	return fmt.Sprintf("Phone => %s, Name => %s, Address => %s", m["phone"], m["name"], m["address"])
}