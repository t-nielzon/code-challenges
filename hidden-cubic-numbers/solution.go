package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func IsSumOfCubes(s string) string {
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
	sum := 0
	for _, g := range groups {
		n, _ := strconv.Atoi(g)
		c := 0
		for _, d := range g {
			v := int(d - '0')
			c += v * v * v
		}
		if c == n {
			cubics = append(cubics, n)
			sum += n
		}
	}

	if len(cubics) == 0 {
		return "Unlucky"
	}

	parts := make([]string, 0, len(cubics)+1)
	for _, n := range cubics {
		parts = append(parts, strconv.Itoa(n))
	}
	parts = append(parts, strconv.Itoa(sum))
	return fmt.Sprintf("%s Lucky", strings.Join(parts, " "))
}