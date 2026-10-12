package main

func IsValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < 10; i++ {
		var digit int
		c := isbn[i]

		if i == 9 {
			if c == 'X' {
				digit = 10
			} else if c >= '0' && c <= '9' {
				digit = int(c - '0')
			} else {
				return false
			}
		} else {
			if c < '0' || c > '9' {
				return false
			}
			digit = int(c - '0')
		}

		sum += digit * (i + 1)
	}

	return sum%11 == 0
}