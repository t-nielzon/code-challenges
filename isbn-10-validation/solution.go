package kata

func ValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i, c := range isbn {
		var value int
		switch {
		case c >= '0' && c <= '9':
			value = int(c - '0')
		case c == 'X' && i == 9:
			// 'X' is only allowed in the final position, where it counts as 10
			value = 10
		default:
			return false
		}
		sum += value * (i + 1)
	}

	return sum%11 == 0
}