package kata

import (
	"strconv"
	"strings"
)

func SimpleAssembler(program []string) map[string]int {
	registers := map[string]int{}

	value := func(s string) int {
		if n, err := strconv.Atoi(s); err == nil {
			return n
		}
		return registers[s]
	}

	for i := 0; i < len(program); {
		parts := strings.Fields(program[i])
		switch parts[0] {
		case "mov":
			registers[parts[1]] = value(parts[2])
			i++
		case "inc":
			registers[parts[1]]++
			i++
		case "dec":
			registers[parts[1]]--
			i++
		case "jnz":
			if value(parts[1]) != 0 {
				i += value(parts[2])
			} else {
				i++
			}
		}
	}

	return registers
}