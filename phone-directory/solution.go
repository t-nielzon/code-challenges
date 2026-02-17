package kata

import (
	"fmt"
	"regexp"
	"strings"
)

func Phone(dir, num string) string {
	lines := strings.Split(dir, "\n")
	var matches []string

	for _, line := range lines {
		if strings.Contains(line, "+"+num) {
			matches = append(matches, line)
		}
	}

	if len(matches) == 0 {
		return fmt.Sprintf("Error => Not found: %s", num)
	}
	if len(matches) > 1 {
		return fmt.Sprintf("Error => Too many people: %s", num)
	}

	line := matches[0]

	// extract name between < and >
	nameRe := regexp.MustCompile(`<([^>]+)>`)
	nameMatch := nameRe.FindStringSubmatch(line)
	name := ""
	if nameMatch != nil {
		name = nameMatch[1]
	}

	// remove the phone number (with leading +) from the line to get the address
	addr := line
	addr = strings.Replace(addr, "+"+num, "", 1)

	// remove the name tag including angle brackets
	addr = nameRe.ReplaceAllString(addr, "")

	// replace underscores with spaces
	addr = strings.ReplaceAll(addr, "_", " ")

	// replace non-alphanumeric characters (except spaces, dots, and dashes) with spaces
	addrRe := regexp.MustCompile(`[^a-zA-Z0-9 .\-]`)
	addr = addrRe.ReplaceAllString(addr, " ")

	// collapse multiple spaces
	spaceRe := regexp.MustCompile(`\s+`)
	addr = spaceRe.ReplaceAllString(addr, " ")

	addr = strings.TrimSpace(addr)

	return fmt.Sprintf("Phone => %s, Name => %s, Address => %s", num, name, addr)
}