package main

import "fmt"

func Rgb(r, g, b int) string {
	clamp := func(val int) int {
		if val < 0 {
			return 0
		}
		if val > 255 {
			return 255
		}
		return val
	}
	
	r = clamp(r)
	g = clamp(g)
	b = clamp(b)
	
	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}