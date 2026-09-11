package main

func DecimalToFactorial(n int) string {
	if n == 0 {
		return "0"
	}
	
	digits := []rune{}
	k := 2
	digitChars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	
	for n > 0 {
		digit := n % k
		digits = append([]rune{rune(digitChars[digit])}, digits...)
		n /= k
		k++
	}
	
	digits = append(digits, '0')
	return string(digits)
}

func FactorialToDecimal(s string) int {
	result := 0
	factorials := []int{1}
	
	for i := 1; i < len(s); i++ {
		factorials = append(factorials, factorials[len(factorials)-1]*i)
	}
	
	for i, ch := range s {
		var digit int
		if ch >= '0' && ch <= '9' {
			digit = int(ch - '0')
		} else {
			digit = int(ch - 'A' + 10)
		}
		result += digit * factorials[len(s)-1-i]
	}
	
	return result
}