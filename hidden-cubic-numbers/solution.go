package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func HiddenCubic(s string) string {
	var groups []string
	i := 0
	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			j := i
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				j++
			}
			run := s[i:j]
			for k := 0; k < len(run); k += 3 {
				end := k + 3
				if end > len(run) {
					end = len(run)
				}
				groups = append(groups, run[k:end])
			}
			i = j
		} else {
			i++
		}
	}

	var cubics []int
	var parts []string
	sum := 0
	for _, g := range groups {
		n, _ := strconv.Atoi(g)
		cubeSum := 0
		for _, c := range g {
			d := int(c - '0')
			cubeSum += d * d * d
		}
		if cubeSum == n {
			cubics = append(cubics, n)
			parts = append(parts, fmt.Sprintf("%d", n))
			sum += n
		}
	}

	if len(cubics) == 0 {
		return "Unlucky"
	}
	parts = append(parts, fmt.Sprintf("%d", sum))
	return strings.Join(parts, " ") + " Lucky"
}