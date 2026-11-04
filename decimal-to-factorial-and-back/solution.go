package main

func ToFactorial(decimal int) string {
	if decimal == 0 {
		return "0"
	}

	digits := []rune{}
	n := decimal

	for i := 1; n > 0; i++ {
		digit := n % i
		digits = append(digits, digitToChar(digit))
		n = n / i
	}

	// reverse digits
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}

func FromFactorial(factorial string) int {
	if factorial == "0" {
		return 0
	}

	result := 0
	power := 1

	for i := len(factorial) - 1; i >= 0; i-- {
		digit := charToDigit(rune(factorial[i]))
		result += digit * power
		power *= (len(factorial) - i)
	}

	return result
}

func digitToChar(digit int) rune {
	if digit < 10 {
		return rune('0' + digit)
	}
	return rune('A' + digit - 10)
}

func charToDigit(ch rune) int {
	if ch >= '0' && ch <= '9' {
		return int(ch - '0')
	}
	return int(ch - 'A' + 10)
}