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

	for pc := 0; pc < len(program); pc++ {
		parts := strings.Fields(program[pc])
		switch parts[0] {
		case "mov":
			regs[parts[1]] = resolve(parts[2])
		case "inc":
			regs[parts[1]]++
		case "dec":
			regs[parts[1]]--
		case "jnz":
			if resolve(parts[1]) != 0 {
				pc += resolve(parts[2]) - 1
			}
		}
	}

	return regs
}