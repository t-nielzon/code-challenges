package kata

import (
	"regexp"
	"strings"
)

func Catalog(s, article string) string {
	prodRe := regexp.MustCompile(`<prod><name>(.*?)</name><prx>(.*?)</prx><qty>(.*?)</qty></prod>`)
	wordRe := regexp.MustCompile(`\b` + regexp.QuoteMeta(article) + `\b`)

	var lines []string
	for _, m := range prodRe.FindAllStringSubmatch(s, -1) {
		name, prx, qty := m[1], m[2], m[3]
		if wordRe.MatchString(name) {
			lines = append(lines, name+" > prx: $"+prx+" qty: "+qty)
		}
	}
	if len(lines) == 0 {
		return "Nothing"
	}
	return strings.Join(lines, "\n")
}