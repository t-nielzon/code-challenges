package kata

import (
	"strings"
	"unicode"
)

func PlayPass(s string, n int) string {
	var buf []byte
	for i, ch := range s {
		var c byte
		switch {
		case unicode.IsLetter(ch):
			shifted := (int(unicode.ToUpper(ch)-'A') + n) % 26
			if i%2 == 0 {
				c = byte('A' + shifted)
			} else {
				c = byte('a' + shifted)
			}
		case unicode.IsDigit(ch):
			c = byte('0' + (9 - (ch - '0')))
		default:
			c = byte(ch)
		}
		buf = append(buf, c)
	}
	// reverse
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return strings.TrimRight(string(buf), "\x00")
}