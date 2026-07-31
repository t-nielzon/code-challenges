package main

func IsValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < 10; i++ {
		char := isbn[i]
		position := i + 1

		var digit int
		if char == 'X' {
			if i != 9 {
				return false
			}
			digit = 10
		} else if char >= '0' && char <= '9' {
			digit = int(char - '0')
		} else {
			return false
		}

		sum += digit * position
	}

	return sum%11 == 0
}