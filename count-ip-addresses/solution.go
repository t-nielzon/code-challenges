package kata

import (
	"fmt"
	"strings"
)

func IpsBetween(start, end string) int {
	return ipToInt(end) - ipToInt(start)
}

func ipToInt(ip string) int {
	parts := strings.Split(ip, ".")
	var a, b, c, d int
	fmt.Sscanf(parts[0], "%d", &a)
	fmt.Sscanf(parts[1], "%d", &b)
	fmt.Sscanf(parts[2], "%d", &c)
	fmt.Sscanf(parts[3], "%d", &d)
	return a<<24 | b<<16 | c<<8 | d
}