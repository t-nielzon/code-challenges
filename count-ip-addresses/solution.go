package kata

import (
	"fmt"
)

func IpsBetween(start, end string) int {
	return ipToInt(end) - ipToInt(start)
}

func ipToInt(ip string) int {
	var a, b, c, d int
	fmt.Sscanf(ip, "%d.%d.%d.%d", &a, &b, &c, &d)
	return a<<24 | b<<16 | c<<8 | d
}