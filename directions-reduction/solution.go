package main

func DirReduc(arr []string) []string {
	stack := []string{}
	
	for _, dir := range arr {
		if len(stack) > 0 && isOpposite(stack[len(stack)-1], dir) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, dir)
		}
	}
	
	return stack
}

func isOpposite(a, b string) bool {
	return (a == "NORTH" && b == "SOUTH") ||
		(a == "SOUTH" && b == "NORTH") ||
		(a == "WEST" && b == "EAST") ||
		(a == "EAST" && b == "WEST")
}