package kata

import (
	"regexp"
	"strings"
)

var prodRe = regexp.MustCompile(`<name>(.*?)</name><prx>(.*?)</prx><qty>(.*?)</qty>`)

func Catalog(s, article string) string {
	matches := prodRe.FindAllStringSubmatch(s, -1)
	var lines []string
	for _, m := range matches {
		name, prx, qty := m[1], m[2], m[3]
		if name == article || strings.HasSuffix(name, " "+article) {
			lines = append(lines, name+" > prx: $"+prx+" qty: "+qty)
		}
	}
	if len(lines) == 0 {
		return "Nothing"
	}
	return strings.Join(lines, "\n")
}