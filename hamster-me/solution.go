package kata

import (
	"strconv"
	"strings"
)

func HamsterMe(code, message string) string {
	var isCode [26]bool
	for _, c := range code {
		isCode[c-'a'] = true
	}

	header := make([]int, 26)
	num := make([]int, 26)
	for i := 0; i < 26; i++ {
		for d := 0; d < 26; d++ {
			p := ((i-d)%26 + 26) % 26
			if isCode[p] {
				header[i] = p
				num[i] = d + 1
				break
			}
		}
	}

	var sb strings.Builder
	for _, ch := range message {
		i := int(ch - 'a')
		sb.WriteByte(byte('a' + header[i]))
		sb.WriteString(strconv.Itoa(num[i]))
	}
	return sb.String()
}