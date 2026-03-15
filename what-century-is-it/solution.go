package kata

import (
	"fmt"
	"strconv"
)

func WhatCentury(year string) string {
	y, _ := strconv.Atoi(year)
	century := (y + 99) / 100

	suffix := "th"
	if century%100 < 11 || century%100 > 13 {
		switch century % 10 {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		case 3:
			suffix = "rd"
		}
	}

	return fmt.Sprintf("%d%s", century, suffix)
}