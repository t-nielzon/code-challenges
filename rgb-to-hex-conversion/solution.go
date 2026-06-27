package kata

import "fmt"

func Rgb(r, g, b int) string {
	return fmt.Sprintf("%02X%02X%02X", clamp(r), clamp(g), clamp(b))
}

func clamp(v int) int {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return v
}