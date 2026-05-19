package kata

func ValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}
	sum := 0
	for i := 0; i < 10; i++ {
		c := isbn[i]
		var v int
		if c >= '0' && c <= '9' {
			v = int(c - '0')
		} else if c == 'X' && i == 9 {
			v = 10
		} else {
			return false
		}
		sum += v * (i + 1)
	}
	return sum%11 == 0
}