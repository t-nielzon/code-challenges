package kata

import "strconv"

func WhatCentury(year string) string {
	y, _ := strconv.Atoi(year)
	c := (y + 99) / 100
	suffix := "th"
	if c%100 < 11 || c%100 > 13 {
		switch c % 10 {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		case 3:
			suffix = "rd"
		}
	}
	return strconv.Itoa(c) + suffix
}