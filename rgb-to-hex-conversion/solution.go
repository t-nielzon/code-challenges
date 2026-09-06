package main

import (
	"fmt"
)

func rgb(r, g, b int) string {
	clamp := func(v int) int {
		if v < 0 {
			return 0
		}
		if v > 255 {
			return 255
		}
		return v
	}
	
	return fmt.Sprintf("%02X%02X%02X", clamp(r), clamp(g), clamp(b))
}