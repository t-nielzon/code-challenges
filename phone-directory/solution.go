package kata

import (
	"fmt"
	"regexp"
	"strings"
)

func phone(strng string, num string) string {
	lines := strings.Split(strng, "\n")
	phoneRe := regexp.MustCompile(`\+\d{1,2}-\d{3}-\d{3}-\d{4}`)
	nameRe := regexp.MustCompile(`<([^>]+)>`)

	var matches []string
	target := "+" + num
	for _, line := range lines {
		if !strings.Contains(line, target) {
			continue
		}
		// confirm it's a full phone match, not partial
		phones := phoneRe.FindAllString(line, -1)
		found := false
		for _, p := range phones {
			if p == target {
				found = true
				break
			}
		}
		if !found {
			continue
		}
		matches = append(matches, line)
	}

	if len(matches) == 0 {
		return fmt.Sprintf("Error => Not found: %s", num)
	}
	if len(matches) > 1 {
		return fmt.Sprintf("Error => Too many people: %s", num)
	}

	line := matches[0]
	nameMatch := nameRe.FindStringSubmatch(line)
	name := ""
	if len(nameMatch) > 1 {
		name = nameMatch[1]
	}

	// remove phone and name from line to get address
	addr := phoneRe.ReplaceAllString(line, "")
	addr = nameRe.ReplaceAllString(addr, "")

	// replace non-alphanumeric (except space, dot, comma maybe) with space
	// keep letters, digits, spaces, and basic punctuation? Looking at example:
	// "/+1-541-754-3010 156 Alphand_St. <J Steeve>" -> "156 Alphand St."
	// so underscore becomes space, slash removed
	var b strings.Builder
	for _, r := range addr {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '.' || r == ' ' {
			b.WriteRune(r)
		} else {
			b.WriteRune(' ')
		}
	}
	addr = b.String()

	// collapse multiple spaces
	spaceRe := regexp.MustCompile(`\s+`)
	addr = spaceRe.ReplaceAllString(addr, " ")
	addr = strings.TrimSpace(addr)

	return fmt.Sprintf("Phone => %s, Name => %s, Address => %s", num, name, addr)
}