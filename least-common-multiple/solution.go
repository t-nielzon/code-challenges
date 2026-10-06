package main

func Lcm(numbers []int) int {
	if len(numbers) == 0 {
		return 1
	}

	for _, num := range numbers {
		if num == 0 {
			return 0
		}
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}