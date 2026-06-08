package kata

import "strings"

func PlayPass(s string, n int) string {
	runes := []rune(s)
	out := make([]rune, len(runes))

	for i, c := range runes {
		switch {
		case c >= 'A' && c <= 'Z':
			c = 'A' + (c-'A'+rune(n))%26
		case c >= '0' && c <= '9':
			c = '0' + (9 - (c - '0'))
		}

		if i%2 == 1 {
			c = []rune(strings.ToLower(string(c)))[0]
		} else {
			c = []rune(strings.ToUpper(string(c)))[0]
		}

		out[len(runes)-1-i] = c
	}

	return string(out)
}