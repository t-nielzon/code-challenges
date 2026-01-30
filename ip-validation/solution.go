package kata

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	parts := strings.Split(ip, ".")
	
	if len(parts) != 4 {
		return false
	}
	
	for _, part := range parts {
		if len(part) == 0 {
			return false
		}
		
		// Check for leading zeros (invalid unless the octet is just "0")
		if len(part) > 1 && part[0] == '0' {
			return false
		}
		
		// Check that all characters are digits
		for _, c := range part {
			if c < '0' || c > '9' {
				return false
			}
		}
		
		// Parse the number and check range
		num, err := strconv.Atoi(part)
		if err != nil {
			return false
		}
		
		if num < 0 || num > 255 {
			return false
		}
	}
	
	return true
}