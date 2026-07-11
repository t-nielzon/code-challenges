package main

import (
	"fmt"
	"strconv"
)

func WhatCentury(year string) string {
	y, _ := strconv.Atoi(year)
	century := (y - 1) / 100 + 1
	
	suffix := "th"
	lastDigit := century % 10
	lastTwoDigits := century % 100
	
	if lastTwoDigits != 11 && lastDigit == 1 {
		suffix = "st"
	} else if lastTwoDigits != 12 && lastDigit == 2 {
		suffix = "nd"
	} else if lastTwoDigits != 13 && lastDigit == 3 {
		suffix = "rd"
	}
	
	return fmt.Sprintf("%d%s", century, suffix)
}