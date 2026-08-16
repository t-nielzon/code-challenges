// solution.go
package kata

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(numbers ...int) int {
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
		result = result * numbers[i] / gcd(result, numbers[i])
	}
	
	return result
}