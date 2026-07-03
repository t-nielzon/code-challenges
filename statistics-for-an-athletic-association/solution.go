package kata

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Stat(strg string) string {
	if strings.TrimSpace(strg) == "" {
		return ""
	}

	parts := strings.Split(strg, ",")
	secs := make([]int, 0, len(parts))
	for _, p := range parts {
		fields := strings.Split(strings.TrimSpace(p), "|")
		h, _ := strconv.Atoi(strings.TrimSpace(fields[0]))
		m, _ := strconv.Atoi(strings.TrimSpace(fields[1]))
		s, _ := strconv.Atoi(strings.TrimSpace(fields[2]))
		secs = append(secs, h*3600+m*60+s)
	}

	sort.Ints(secs)
	n := len(secs)

	rng := secs[n-1] - secs[0]

	sum := 0
	for _, v := range secs {
		sum += v
	}
	avg := sum / n // truncated

	var median int
	if n%2 == 1 {
		median = secs[n/2]
	} else {
		median = (secs[n/2-1] + secs[n/2]) / 2 // truncated
	}

	return fmt.Sprintf("Range: %s Average: %s Median: %s",
		format(rng), format(avg), format(median))
}

func format(total int) string {
	h := total / 3600
	m := (total % 3600) / 60
	s := total % 60
	return fmt.Sprintf("%02d|%02d|%02d", h, m, s)
}