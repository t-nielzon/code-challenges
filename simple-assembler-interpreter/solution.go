package kata

import (
	"strconv"
	"strings"
)

func SimpleAssembler(program []string) map[string]int {
	registers := make(map[string]int)

	value := func(token string) int {
		if n, err := strconv.Atoi(token); err == nil {
			return n
		}
		return registers[token]
	}

	for i := 0; i < len(program); i++ {
		parts := strings.Fields(program[i])
		switch parts[0] {
		case "mov":
			registers[parts[1]] = value(parts[2])
		case "inc":
			registers[parts[1]]++
		case "dec":
			registers[parts[1]]--
		case "jnz":
			if value(parts[1]) != 0 {
				i += value(parts[2]) - 1
			}
		}
	}

	return registers
}