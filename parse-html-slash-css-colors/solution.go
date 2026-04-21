package kata

import (
	"strconv"
	"strings"
)

type Color struct{ R, G, B byte }

func Parse(s string) Color {
	if !strings.HasPrefix(s, "#") {
		s = PresetColors[strings.ToLower(s)]
	}
	hex := s[1:]
	if len(hex) == 3 {
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	}
	n, _ := strconv.ParseUint(hex, 16, 32)
	return Color{byte(n >> 16), byte(n >> 8), byte(n)}
}