package kata

func DirReduc(arr []string) []string {
	opposites := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"EAST":  "WEST",
		"WEST":  "EAST",
	}

	stack := []string{}

	for _, dir := range arr {
		if len(stack) > 0 && opposites[dir] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, dir)
		}
	}

	return stack
}