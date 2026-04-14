package kata

func IsValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i, ch := range isbn {
		var val int
		if ch >= '0' && ch <= '9' {
			val = int(ch - '0')
		} else if ch == 'X' && i == 9 {
			val = 10
		} else {
			return false
		}
		sum += val * (i + 1)
	}

	return sum%11 == 0
}