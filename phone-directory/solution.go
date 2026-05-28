package kata

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	nameRe    = regexp.MustCompile(`<(.+?)>`)
	cleanRe   = regexp.MustCompile(`[^a-zA-Z0-9 .]`)
	spacesRe  = regexp.MustCompile(`\s+`)
	dotSpaceR = regexp.MustCompile(`\s+\.`)
)

func phone(strng string, num string) string {
	var matches []string
	for _, line := range strings.Split(strng, "\n") {
		if strings.Contains(line, "+"+num+" ") ||
			strings.Contains(line, "+"+num+"\t") ||
			strings.HasSuffix(strings.TrimRight(line, "!;., \t"), "+"+num) ||
			strings.Contains(line, "+"+num+"!") ||
			strings.Contains(line, "+"+num+";") ||
			strings.Contains(line, "+"+num+",") {
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

	name := ""
	if m := nameRe.FindStringSubmatch(line); m != nil {
		name = m[1]
	}

	// strip the phone number and the <name> block, leaving only the address.
	addr := line
	addr = strings.Replace(addr, "+"+num, " ", 1)
	addr = nameRe.ReplaceAllString(addr, " ")

	addr = cleanRe.ReplaceAllString(addr, " ")
	addr = spacesRe.ReplaceAllString(addr, " ")
	addr = dotSpaceR.ReplaceAllString(addr, ".")
	addr = strings.TrimSpace(addr)

	return fmt.Sprintf("Phone => %s, Name => %s, Address => %s", num, name, addr)
}