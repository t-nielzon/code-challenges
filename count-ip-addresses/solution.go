package kata

import (
	"strconv"
	"strings"
)

func ipToUint(ip string) uint32 {
	parts := strings.Split(ip, ".")
	var n uint32
	for _, p := range parts {
		v, _ := strconv.Atoi(p)
		n = n<<8 | uint32(v)
	}
	return n
}

func IpsBetween(start string, end string) uint32 {
	return ipToUint(end) - ipToUint(start)
}