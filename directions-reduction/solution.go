package kata

func dirReduc(arr []string) []string {
	opposite := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"EAST":  "WEST",
		"WEST":  "EAST",
	}

	stack := []string{}
	for _, dir := range arr {
		if len(stack) > 0 && stack[len(stack)-1] == opposite[dir] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, dir)
		}
	}

	return stack
}