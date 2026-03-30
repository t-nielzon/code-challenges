package kata

func DirReduc(arr []string) []string {
	opposite := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"EAST":  "WEST",
		"WEST":  "EAST",
	}

	stack := make([]string, 0, len(arr))
	for _, dir := range arr {
		if len(stack) > 0 && opposite[stack[len(stack)-1]] == dir {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, dir)
		}
	}
	return stack
}