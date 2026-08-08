package main

func DecimalToFactorial(n int) string {
	if n == 0 {
		return "0"
	}
	
	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result []rune
	
	position := 1
	for n > 0 {
		digit := n % (position + 1)
		result = append(result, rune(digits[digit]))
		n = n / (position + 1)
		position++
	}
	
	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	
	// Add trailing 0 for 0! position
	result = append(result, '0')
	
	return string(result)
}

func FactorialToDecimal(s string) int {
	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := 0
	
	for i, ch := range s {
		digitValue := 0
		for j, d := range digits {
			if d == ch {
				digitValue = j
				break
			}
		}
		
		base := len(s) - 1 - i
		baseFactorial := 1
		for j := 1; j <= base; j++ {
			baseFactorial *= j
		}
		
		result += digitValue * baseFactorial
	}
	
	return result
}