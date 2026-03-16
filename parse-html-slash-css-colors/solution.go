package kata

import (
	"strconv"
	"strings"
)

func Parse(color string) Color {
	if color[0] == '#' {
		hex := color[1:]
		if len(hex) == 3 {
			hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
		}
		r, _ := strconv.ParseUint(hex[0:2], 16, 8)
		g, _ := strconv.ParseUint(hex[2:4], 16, 8)
		b, _ := strconv.ParseUint(hex[4:6], 16, 8)
		return Color{byte(r), byte(g), byte(b)}
	}
	return Parse(PresetColors[strings.ToLower(color)])
}