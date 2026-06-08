package kata

import (
	"strconv"
	"strings"
)

func isValidIP(ip string) bool {
	octets := strings.Split(ip, ".")
	if len(octets) != 4 {
		return false
	}
	for _, octet := range octets {
		// reject empty segments and leading zeros (e.g. "01", "007")
		if len(octet) == 0 || (len(octet) > 1 && octet[0] == '0') {
			return false
		}
		n, err := strconv.Atoi(octet)
		if err != nil || n < 0 || n > 255 {
			return false
		}
	}
	return true
}