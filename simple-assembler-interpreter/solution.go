package kata

import (
	"strconv"
	"strings"
)

func SimpleAssembler(program []string) map[string]int {
	regs := map[string]int{}

	resolve := func(s string) int {
		if v, err := strconv.Atoi(s); err == nil {
			return v
		}
		return regs[s]
	}

	for i := 0; i < len(program); {
		parts := strings.Fields(program[i])
		switch parts[0] {
		case "mov":
			regs[parts[1]] = resolve(parts[2])
		case "inc":
			regs[parts[1]]++
		case "dec":
			regs[parts[1]]--
		case "jnz":
			if resolve(parts[1]) != 0 {
				i += resolve(parts[2])
				continue
			}
		}
		i++
	}

	return regs
}