package kata

import (
	"strconv"
	"strings"
)

func IsValidIP(str string) bool {
	parts := strings.Split(str, ".")
	if len(parts) != 4 {
		return false
	}
	for _, p := range parts {
		if len(p) == 0 || len(p) > 3 {
			return false
		}
		if len(p) > 1 && p[0] == '0' {
			return false
		}
		for _, c := range p {
			if c < '0' || c > '9' {
				return false
			}
		}
		n, err := strconv.Atoi(p)
		if err != nil || n < 0 || n > 255 {
			return false
		}
	}
	return true
}