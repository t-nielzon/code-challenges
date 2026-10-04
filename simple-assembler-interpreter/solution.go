package main

import (
	"strconv"
	"strings"
)

func SimpleAssemblerInterpreter(program []string) map[string]int {
	registers := make(map[string]int)

	getValue := func(s string) int {
		// check if it's a register (single lowercase letter)
		if len(s) == 1 && s[0] >= 'a' && s[0] <= 'z' {
			return registers[s]
		}
		// otherwise parse as integer constant
		val, _ := strconv.Atoi(s)
		return val
	}

	for i := 0; i < len(program); i++ {
		parts := strings.Fields(program[i])
		cmd := parts[0]

		switch cmd {
		case "mov":
			registers[parts[1]] = getValue(parts[2])
		case "inc":
			registers[parts[1]]++
		case "dec":
			registers[parts[1]]--
		case "jnz":
			if getValue(parts[1]) != 0 {
				// jump y steps; adjust by -1 since loop increments i by 1
				i += getValue(parts[2]) - 1
			}
		}
	}

	return registers
}