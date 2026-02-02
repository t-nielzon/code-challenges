package kata

func LCM(args []int) int {
	if len(args) == 0 {
		return 1
	}

	result := args[0]
	for i := 1; i < len(args); i++ {
		result = lcm(result, args[i])
		if result == 0 {
			return 0
		}
	}
	return result
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return (a / gcd(a, b)) * b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}