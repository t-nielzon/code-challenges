package main

import "strings"

func IsValidIP(ip string) bool {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}
	
	for _, part := range parts {
		if len(part) == 0 {
			return false
		}
		
		if len(part) > 1 && part[0] == '0' {
			return false
		}
		
		for _, ch := range part {
			if ch < '0' || ch > '9' {
				return false
			}
		}
		
		num := 0
		for _, ch := range part {
			num = num*10 + int(ch-'0')
		}
		
		if num > 255 {
			return false
		}
	}
	
	return true
}